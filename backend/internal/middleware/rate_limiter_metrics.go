package middleware

// Snapshot 返回当前所有调用方桶的副本（诊断用）。
func (rl *RateLimiter) Snapshot() map[uint]*bucket {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	out := make(map[uint]*bucket, len(rl.buckets))
	for k, v := range rl.buckets {
		out[k] = v
	}
	return out
}

// BucketCount 返回当前限流桶数量。
func (rl *RateLimiter) BucketCount() int {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	return len(rl.buckets)
}
