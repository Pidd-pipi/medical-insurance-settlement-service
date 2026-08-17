# BUG_REPRO

## Bug 是什么
日终对账查询使用 PostgreSQL 专属的 `settled_at::date` 写法，在 SQLite 上直接报错；同时 reversed 状态未被失败数统计覆盖。

## 如何触发
在 backend 目录运行：

```bash
go test ./internal/service -run TestReconciliationService_Daily_ClassifiesReversed -count=1
```

测试会创建一条当日 reversed 结算单并生成对账。

## 错误信息
```
SQL logic error: unrecognized token: ":"
```
