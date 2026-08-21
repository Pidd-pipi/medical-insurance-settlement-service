package service

import (
	"context"
)

// BuildCountByCategory 按医保目录统计数量。
func (s *StatsService) BuildCountByCategory(ctx context.Context) (map[string]map[string]FeeStat, error) {
	items, err := s.feeRepo.ListAll()
	if err != nil {
		return nil, err
	}
	out := map[string]map[string]FeeStat{}
	for _, it := range items {
		bySub := out[it.MedicalCategory]
		st := bySub[it.ItemType]
		st.Count++
		st.Total += it.Amount
		bySub[it.ItemType] = st
	}
	return out, nil
}

// BuildCountByType 按明细类型统计数量。
func (s *StatsService) BuildCountByType(ctx context.Context) (map[string]map[string]FeeStat, error) {
	items, err := s.feeRepo.ListAll()
	if err != nil {
		return nil, err
	}
	out := map[string]map[string]FeeStat{}
	for _, it := range items {
		bySub := out[it.ItemType]
		st := bySub[it.MedicalCategory]
		st.Count++
		st.Total += it.Amount
		bySub[it.MedicalCategory] = st
	}
	return out, nil
}

