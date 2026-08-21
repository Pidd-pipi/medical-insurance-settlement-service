# BUG_REPRO

## Bug 是什么
参保人核验接口查询一个不存在的参保人时返回 500，而不是预期的 404。错误链路在仓储、服务、接口三层同时错位：仓储用 `%v` 断掉哨兵错误、服务判断了错误的哨兵、接口把错误统一包成内部错误。

## 如何触发
在 backend 目录运行：

```bash
go test ./internal/service -run '^TestInsuredVerifyMissingReturns404$' -count=1
```

## 错误信息
```
--- FAIL: TestInsuredVerifyMissingReturns404
    insurance_service_test.go:22: expected 404 AppError, got verify insured: find insured by id_card and medical_card: record not found
```
