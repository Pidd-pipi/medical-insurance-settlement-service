package service

import (
	"context"
	"testing"

	"github.com/blueship581/gbinsureapi/internal/constants"
	"github.com/blueship581/gbinsureapi/internal/model"
	"github.com/blueship581/gbinsureapi/internal/repository"
)

func TestApiClientService_UpdateStatusRejectsInvalid(t *testing.T) {
	db := newTestDB(t)
	client := model.ApiClient{Name: "测试HIS", ClientType: constants.ClientTypeHIS, APIKeyHash: "h", Role: "settlement", Status: constants.ClientActive, RateLimitQPS: 10}
	if err := db.Create(&client).Error; err != nil {
		t.Fatal(err)
	}
	svc := NewApiClientService(repository.NewApiClientRepository(db), "api-secret", "jwt-secret", 24, testLogger())
	if err := svc.UpdateStatus(context.Background(), client.ID, "bogus"); err == nil {
		t.Fatal("expected invalid client status error")
	}
}
