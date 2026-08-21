package service

import (
	"context"
	"testing"

	"github.com/blueship581/gbinsureapi/internal/constants"
	"github.com/blueship581/gbinsureapi/internal/model"
	"github.com/blueship581/gbinsureapi/internal/repository"
)

func seedStatsItems(t *testing.T, feeRepo *repository.FeeItemRepository) {
	t.Helper()
	items := []model.FeeItem{
		{BatchID: 1, ItemCode: "DRUG01", ItemName: "阿莫西林", ItemType: constants.FeeItemDrug, Amount: 50, MedicalCategory: constants.MedicalCategoryClassA},
		{BatchID: 1, ItemCode: "EXAM01", ItemName: "血常规", ItemType: constants.FeeItemExam, Amount: 40, MedicalCategory: constants.MedicalCategoryClassA},
		{BatchID: 1, ItemCode: "BED01", ItemName: "床位费", ItemType: constants.FeeItemTreatment, Amount: 100, MedicalCategory: constants.MedicalCategoryClassB},
	}
	if err := feeRepo.CreateBatch(items); err != nil {
		t.Fatal(err)
	}
}

func TestStatsBuildNoPanic(t *testing.T) {
	db := newTestDB(t)
	feeRepo := repository.NewFeeItemRepository(db)
	seedStatsItems(t, feeRepo)
	out, err := NewStatsService(feeRepo, testLogger()).Build(context.Background())
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	if got := out[constants.MedicalCategoryClassA][constants.FeeItemDrug]; got.Count != 1 || got.Total != 50 {
		t.Fatalf("class_a drug = %+v", got)
	}
}

func TestStatsBuildByTypeFirstNoPanic(t *testing.T) {
	db := newTestDB(t)
	feeRepo := repository.NewFeeItemRepository(db)
	seedStatsItems(t, feeRepo)
	out, err := NewStatsService(feeRepo, testLogger()).BuildByTypeFirst(context.Background())
	if err != nil {
		t.Fatalf("BuildByTypeFirst() error = %v", err)
	}
	if got := out[constants.FeeItemDrug][constants.MedicalCategoryClassA]; got.Count != 1 || got.Total != 50 {
		t.Fatalf("drug class_a = %+v", got)
	}
}

func TestStatsBuildCountByCategoryNoPanic(t *testing.T) {
	db := newTestDB(t)
	feeRepo := repository.NewFeeItemRepository(db)
	seedStatsItems(t, feeRepo)
	out, err := NewStatsService(feeRepo, testLogger()).BuildCountByCategory(context.Background())
	if err != nil {
		t.Fatalf("BuildCountByCategory() error = %v", err)
	}
	if got := out[constants.MedicalCategoryClassB][constants.FeeItemTreatment]; got.Count != 1 {
		t.Fatalf("class_b treatment = %+v", got)
	}
}

func TestStatsBuildCountByTypeNoPanic(t *testing.T) {
	db := newTestDB(t)
	feeRepo := repository.NewFeeItemRepository(db)
	seedStatsItems(t, feeRepo)
	out, err := NewStatsService(feeRepo, testLogger()).BuildCountByType(context.Background())
	if err != nil {
		t.Fatalf("BuildCountByType() error = %v", err)
	}
	if got := out[constants.FeeItemExam][constants.MedicalCategoryClassA]; got.Count != 1 {
		t.Fatalf("exam class_a = %+v", got)
	}
}
