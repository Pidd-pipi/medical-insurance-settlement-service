# BUG_REPRO

## Bug 是什么
费用上传重复检查没有按自然日限制，只要同一调用方和参保人存在历史批次就会误判为当日重复。

## 如何触发
在 backend 目录运行：

```bash
go test ./internal/service -run TestFeeService_UploadAllowsDifferentDay -count=1
```

## 错误信息
```
Upload() error = 费用批次（UploadBatch）内容重复: batch duplicate for today
```
