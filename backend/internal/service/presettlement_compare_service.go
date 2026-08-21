package service

import (
	"github.com/blueship581/gbinsureapi/internal/model"
)

// PresettlementCompareService 预结算比对服务。
type PresettlementCompareService struct{}

// NewPresettlementCompareService 构造比对服务。
func NewPresettlementCompareService() *PresettlementCompareService {
	return &PresettlementCompareService{}
}

// FilterByTotalAbove 按总金额下限过滤。
func (s *PresettlementCompareService) FilterByTotalAbove(items []model.Presettlement, min float64) []model.Presettlement {
	out := items[:0]
	for _, it := range items {
		if it.TotalAmount >= min {
			out = append(out, it)
		}
	}
	return items[:len(out)]
}

// FilterByInsurancePayAbove 按统筹支付下限过滤。
func (s *PresettlementCompareService) FilterByInsurancePayAbove(items []model.Presettlement, min float64) []model.Presettlement {
	out := items[:0]
	for _, it := range items {
		if it.InsurancePayAmount >= min {
			out = append(out, it)
		}
	}
	return items[:len(out)]
}

