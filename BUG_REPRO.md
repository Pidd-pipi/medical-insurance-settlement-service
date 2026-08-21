# BUG_REPRO

## Bug 是什么
登录接口限流在并发下计数不准，该拦的请求拦不住，并发压测还会报 data race。根因是令牌桶计数更新在锁外执行、诊断快照直接返回内部 map 引用、指标读取无锁。

## 如何触发
在 backend 目录运行：

```bash
go test -race ./internal/middleware -run '^TestRateLimiterConcurrentNoRace$' -count=1
```

## 错误信息
```
WARNING: DATA RACE
Read at ... by goroutine 19:
  ...RateLimiter.allow()
      .../rate_limiter.go:44
Previous write at ... by goroutine 22:
  ...RateLimiter.allow()
      .../rate_limiter.go:49
```
