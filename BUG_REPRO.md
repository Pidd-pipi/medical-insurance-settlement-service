# BUG_REPRO

## Bug 是什么
请求超时或客户端断开后，后台仍会继续写审计日志。根因是审计中间件丢弃了请求 ctx 改用 Background，仓储也不检查 ctx 是否已取消。

## 如何触发
在 backend 目录运行：

```bash
go test ./internal/middleware -run '^TestAuditLogSkipWhenCanceled$' -count=1
```

## 错误信息
```
--- FAIL: TestAuditLogSkipWhenCanceled
    audit_log_test.go:50: audit log written despite canceled ctx: count=1
```
