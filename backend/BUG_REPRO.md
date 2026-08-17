# BUG_REPRO

## Bug 是什么
结算单列表的总数统计带上了 status 过滤，但实际取数查询没有 status 过滤，导致返回结果混入其他状态。

## 如何触发
在 backend 目录运行：

```bash
go test ./internal/service -run TestSettlementService_ListOrdersFiltersStatus -count=1
```

## 错误信息
```
total=1 len=2 status=reversed, want settled only
```
