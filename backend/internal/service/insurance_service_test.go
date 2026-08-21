package service

import (
	"context"
	"errors"
	"testing"

	"github.com/blueship581/gbinsureapi/internal/repository"
	"github.com/blueship581/gbinsureapi/internal/util"
)

func TestInsuredVerifyMissingReturns404(t *testing.T) {
	db := newTestDB(t)
	ctx := context.Background()
	svc := NewInsuranceService(repository.NewInsuredPersonRepository(db), testLogger())
	_, err := svc.Verify(ctx, "000000000000000000", "M000000000000000000")
	if err == nil {
		t.Fatal("expected error for missing insured")
	}
	var appErr *util.AppError
	if !errors.As(err, &appErr) || appErr.HTTPStatus != 404 {
		t.Fatalf("expected 404 AppError, got %v", err)
	}
}

func TestInsuredGetByIDMissingReturns404(t *testing.T) {
	db := newTestDB(t)
	ctx := context.Background()
	svc := NewInsuranceService(repository.NewInsuredPersonRepository(db), testLogger())
	_, err := svc.GetByID(ctx, 999999)
	if err == nil {
		t.Fatal("expected error for missing insured")
	}
	var appErr *util.AppError
	if !errors.As(err, &appErr) || appErr.HTTPStatus != 404 {
		t.Fatalf("expected 404 AppError, got %v", err)
	}
}
