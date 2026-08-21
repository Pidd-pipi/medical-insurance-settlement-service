package middleware

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestRateLimiterSnapshotIsolated(t *testing.T) {
	rl := NewRateLimiter()
	rl.allow(1, 10)
	snap := rl.Snapshot()
	rl.allow(2, 10)
	if len(snap) != 1 {
		t.Fatalf("snapshot leaked internal map: len=%d, want 1", len(snap))
	}
}

func TestRateLimiterBucketCountConsistent(t *testing.T) {
	rl := NewRateLimiter()
	var wg sync.WaitGroup
	start := make(chan struct{})
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func(id uint) {
			defer wg.Done()
			<-start
			for j := 0; j < 30; j++ {
				rl.allow(id, 50)
			}
		}(uint(i + 1))
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-start
		for j := 0; j < 300; j++ {
			_ = rl.BucketCount()
		}
	}()
	close(start)
	wg.Wait()
	if rl.BucketCount() != 16 {
		t.Fatalf("bucket count = %d, want 16", rl.BucketCount())
	}
}

func TestRateLimiterConcurrentNoRace(t *testing.T) {
	rl := NewRateLimiter()
	const workers = 32
	var wg sync.WaitGroup
	start := make(chan struct{})
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 50; j++ {
				rl.allow(1, 100)
			}
		}()
	}
	close(start)
	wg.Wait()
}

func TestRateLimiterTokenAccounting(t *testing.T) {
	rl := NewRateLimiter()
	const qps = 20
	var allowed int32
	var wg sync.WaitGroup
	start := make(chan struct{})
	for i := 0; i < qps; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			if rl.allow(1, qps) {
				atomic.AddInt32(&allowed, 1)
			}
		}()
	}
	close(start)
	wg.Wait()
	extra := 0
	for i := 0; i < 5; i++ {
		if rl.allow(1, qps) {
			extra++
		}
	}
	if extra > 1 {
		t.Fatalf("token accounting broken: extra allowed=%d, want <=1", extra)
	}
}
