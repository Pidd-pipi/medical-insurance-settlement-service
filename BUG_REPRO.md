# BUG_REPRO

## Bug 是什么
预结算比对功能第二次过滤时会串数据，读到上一次过滤留下的内容。根因是过滤用 `items[:0]` 原地压缩共享底层数组，污染了调用方保留的原始列表。

## 如何触发
在 backend 目录运行：

```bash
go test ./internal/service -run '^TestPresettlementFilterByTotalAboveKeepsSource$' -count=1
```

## 错误信息
```
--- FAIL: TestPresettlementFilterByTotalAboveKeepsSource
    presettlement_compare_service_test.go:19: source slice corrupted: [{ID:2 ...
```
