# BUG_REPRO

## Bug 是什么
新农合医保类型的起付线被错误配置为 300，而不是 200，导致预结算金额计算错误。

## 如何触发
在 backend 目录运行：

```bash
go test ./internal/util -run TestSettlementCalculator_NewRuralPolicy -count=1
```

## 错误信息
```
Deductible = 300, want 200
```
