package service

import (
	"github.com/blueship581/gbinsureapi/internal/model"
)

// FilterBySelfPayAbove 按自费金额下限过滤。
func (s *PresettlementCompareService) FilterBySelfPayAbove(items []model.Presettlement, min float64) []model.Presettlement {
	out := items[:0]
	for _, it := range items {
		if it.SelfPayAmount >= min {
			out = append(out, it)
		}
	}
	return items[:len(out)]
}

// FilterByDeductibleBelow 按起付线上限过滤。
func (s *PresettlementCompareService) FilterByDeductibleBelow(items []model.Presettlement, max float64) []model.Presettlement {
	out := items[:0]
	for _, it := range items {
		if it.Deductible <= max {
			out = append(out, it)
		}
	}
	return items[:len(out)]
}

