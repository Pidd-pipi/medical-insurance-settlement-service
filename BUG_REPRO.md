# BUG_REPRO

## Bug 是什么
费用统计功能一跑就崩溃，日志报 nil map 写入。根因是统计聚合的外层 map 已初始化，但内层 map 在首次写入前未初始化。

## 如何触发
在 backend 目录运行：

```bash
go test ./internal/service -run '^TestStatsBuildNoPanic$' -count=1
```

## 错误信息
```
panic: assignment to entry in nil map
...(*StatsService).Build(...)
    .../stats_service.go:39
```
