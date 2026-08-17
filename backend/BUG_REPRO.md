# BUG_REPRO

## Bug 是什么
调用方状态更新没有校验状态枚举，非法字符串也会写入数据库。

## 如何触发
在 backend 目录运行：

```bash
go test ./internal/service -run TestApiClientService_UpdateStatusRejectsInvalid -count=1
```

## 错误信息
```
expected invalid client status error
```
