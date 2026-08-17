package service

import (
	"context"
	"testing"
	"time"

	"github.com/blueship581/gbinsureapi/internal/constants"
	"github.com/blueship581/gbinsureapi/internal/model"
	"github.com/blueship581/gbinsureapi/internal/repository"
)

func TestFeeService_UploadAllowsDifferentDay(t *testing.T) {
	db := newTestDB(t)
	ctx := context.Background()
	insurance := NewInsuranceService(repository.NewInsuredPersonRepository(db), testLogger())
	batchRepo := repository.NewUploadBatchRepository(db)
	feeRepo := repository.NewFeeItemRepository(db)
	svc := NewFeeService(batchRepo, feeRepo, insurance, testLogger())

	person := model.InsuredPerson{IDCardNo: "110101199001011234", MedicalCardNo: "M110101199001011234", Name: "张三", InsuranceType: constants.InsuranceTypeEmployee, InsuranceStatus: constants.InsuranceStatusActive, PersonalBalance: 3000}
	if err := db.Create(&person).Error; err != nil {
		t.Fatal(err)
	}
	old := model.UploadBatch{BatchNo: "B20260816000001", ClientID: 1, InsuredPersonID: person.ID, UploadStatus: constants.UploadValidated, CreatedAt: time.Now().AddDate(0, 0, -1)}
	if err := db.Create(&old).Error; err != nil {
		t.Fatal(err)
	}

	res, err := svc.Upload(ctx, UploadInput{ClientID: 1, InsuredPersonID: person.ID, Items: []FeeItemInput{{ItemCode: "DRUG01", ItemName: "阿莫西林", ItemType: constants.FeeItemDrug, UnitPrice: 10, Quantity: 1, MedicalCategory: constants.MedicalCategoryClassA}}})
	if err != nil {
		t.Fatalf("Upload() error = %v", err)
	}
	if res.UploadStatus != constants.UploadValidated {
		t.Fatalf("UploadStatus = %s, want validated", res.UploadStatus)
	}
}
