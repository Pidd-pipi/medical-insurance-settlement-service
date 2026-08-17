package service

import (
	"context"
	"testing"
	"time"

	"github.com/blueship581/gbinsureapi/internal/constants"
	"github.com/blueship581/gbinsureapi/internal/model"
	"github.com/blueship581/gbinsureapi/internal/repository"
)

func TestReconciliationService_Daily_ClassifiesReversed(t *testing.T) {
	db := newTestDB(t)
	orderRepo := repository.NewSettlementOrderRepository(db)
	recRepo := repository.NewDailyReconciliationRepository(db)
	svc := NewReconciliationService(orderRepo, recRepo, testLogger())

	now := time.Now()
	order := &model.SettlementOrder{
		SettlementNo: "SETTLE-RECON-001", BatchID: 1, InsuredPersonID: 1,
		PresettlementID: 1, ClientID: 7, Status: constants.SettlementReversed,
		TotalAmount: 100, InsurancePayAmount: 40, SettledAt: &now,
	}
	if err := db.Create(order).Error; err != nil {
		t.Fatalf("seed order: %v", err)
	}

	rec, err := svc.Daily(context.Background(), 7)
	if err != nil {
		t.Fatalf("Daily() error = %v", err)
	}
	if rec.TotalCount != 1 {
		t.Fatalf("TotalCount = %d, want 1", rec.TotalCount)
	}
	if rec.SuccessCount != 0 {
		t.Fatalf("SuccessCount = %d, want 0", rec.SuccessCount)
	}
	if rec.FailCount != 1 {
		t.Fatalf("FailCount = %d, want 1", rec.FailCount)
	}
}
