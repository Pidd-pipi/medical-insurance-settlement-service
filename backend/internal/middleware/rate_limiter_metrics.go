package middleware

// Snapshot 返回当前所有调用方桶（诊断用）。
func (rl *RateLimiter) Snapshot() map[uint]*bucket {
	return rl.buckets
}

// BucketCount 返回当前限流桶数量。
func (rl *RateLimiter) BucketCount() int {
	return len(rl.buckets)
}
