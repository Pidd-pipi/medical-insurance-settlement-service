package repository

import (
	"context"
	"testing"

	"github.com/blueship581/gbinsureapi/internal/model"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func newAuditRepoTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&model.AuditLog{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db
}

func TestAuditLogRepositoryCreateRespectsCtx(t *testing.T) {
	db := newAuditRepoTestDB(t)
	repo := NewAuditLogRepository(db)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := repo.Create(ctx, &model.AuditLog{ClientID: 1, Method: "GET", Path: "/x"})
	if err == nil {
		t.Fatal("expected error for canceled ctx")
	}
	var count int64
	db.Model(&model.AuditLog{}).Count(&count)
	if count != 0 {
		t.Fatalf("audit log written despite canceled ctx: count=%d", count)
	}
}

func TestAuditLogRepositoryBatchCreateRespectsCtx(t *testing.T) {
	db := newAuditRepoTestDB(t)
	repo := NewAuditLogRepository(db)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := repo.BatchCreate(ctx, []model.AuditLog{{ClientID: 1, Method: "GET", Path: "/x"}})
	if err == nil {
		t.Fatal("expected error for canceled ctx")
	}
	var count int64
	db.Model(&model.AuditLog{}).Count(&count)
	if count != 0 {
		t.Fatalf("audit log written despite canceled ctx: count=%d", count)
	}
}
