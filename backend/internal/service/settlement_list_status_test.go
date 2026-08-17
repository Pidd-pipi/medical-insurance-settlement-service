package service

import (
	"context"
	"testing"
	"time"

	"github.com/blueship581/gbinsureapi/internal/constants"
	"github.com/blueship581/gbinsureapi/internal/model"
	"github.com/blueship581/gbinsureapi/internal/repository"
	"github.com/blueship581/gbinsureapi/internal/util"
)

func TestListOrdersFiltersStatus(t *testing.T) {
	db := newTestDB(t)
	now := time.Now()
	if err := db.Create(&model.SettlementOrder{SettlementNo: "S1", BatchID: 1, InsuredPersonID: 1, PresettlementID: 1, ClientID: 7, Status: constants.SettlementSettled, SettledAt: &now}).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Create(&model.SettlementOrder{SettlementNo: "S2", BatchID: 2, InsuredPersonID: 1, PresettlementID: 1, ClientID: 7, Status: constants.SettlementReversed, SettledAt: &now}).Error; err != nil {
		t.Fatal(err)
	}
	svc := NewSettlementService(repository.NewPresettlementRepository(db), repository.NewSettlementOrderRepository(db), repository.NewFeeItemRepository(db), repository.NewUploadBatchRepository(db), NewInsuranceService(repository.NewInsuredPersonRepository(db), testLogger()), util.NewSettlementCalculator(), testLogger())
	items, total, err := svc.ListOrders(context.Background(), 7, constants.SettlementSettled, 1, 10)
	if err != nil {
		t.Fatal(err)
	}
	if total != 1 || len(items) != 1 || items[0].Status != constants.SettlementSettled {
		t.Fatalf("total=%d len=%d status=%s, want settled only", total, len(items), items[0].Status)
	}
}

