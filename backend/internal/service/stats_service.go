package service

import (
	"context"
	"log/slog"

	"github.com/blueship581/gbinsureapi/internal/repository"
)

// StatsService 费用统计服务。
type StatsService struct {
	feeRepo *repository.FeeItemRepository
	log     *slog.Logger
}

// NewStatsService 构造统计服务。
func NewStatsService(feeRepo *repository.FeeItemRepository, log *slog.Logger) *StatsService {
	return &StatsService{feeRepo: feeRepo, log: log}
}

// FeeStat 单个分类统计。
type FeeStat struct {
	Count int     `json:"count"`
	Total float64 `json:"total"`
}

// Build 按医保目录与明细类型聚合。
func (s *StatsService) Build(ctx context.Context) (map[string]map[string]FeeStat, error) {
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

// BuildByTypeFirst 按明细类型与医保目录聚合。
func (s *StatsService) BuildByTypeFirst(ctx context.Context) (map[string]map[string]FeeStat, error) {
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

