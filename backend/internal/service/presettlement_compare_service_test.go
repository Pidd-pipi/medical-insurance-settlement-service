package service

import (
	"testing"

	"github.com/blueship581/gbinsureapi/internal/model"
)

func samplePresettlements() []model.Presettlement {
	return []model.Presettlement{
		{ID: 1, TotalAmount: 100, InsurancePayAmount: 60, SelfPayAmount: 40, Deductible: 600},
		{ID: 2, TotalAmount: 300, InsurancePayAmount: 220, SelfPayAmount: 80, Deductible: 200},
		{ID: 3, TotalAmount: 150, InsurancePayAmount: 90, SelfPayAmount: 60, Deductible: 500},
	}
}

func assertSourceIntact(t *testing.T, items []model.Presettlement) {
	t.Helper()
	if len(items) != 3 || items[0].ID != 1 || items[1].ID != 2 || items[2].ID != 3 {
		t.Fatalf("source slice corrupted: %+v", items)
	}
}

func assertIDs(t *testing.T, got []model.Presettlement, want ...uint) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d (%+v)", len(got), len(want), got)
	}
	for i, id := range want {
		if got[i].ID != id {
			t.Fatalf("got[%d].ID = %d, want %d", i, got[i].ID, id)
		}
	}
}

func TestPresettlementFilterByTotalAboveKeepsSource(t *testing.T) {
	items := samplePresettlements()
	got := NewPresettlementCompareService().FilterByTotalAbove(items, 120)
	assertSourceIntact(t, items)
	assertIDs(t, got, 2, 3)
}

func TestPresettlementFilterByInsurancePayAboveKeepsSource(t *testing.T) {
	items := samplePresettlements()
	got := NewPresettlementCompareService().FilterByInsurancePayAbove(items, 80)
	assertSourceIntact(t, items)
	assertIDs(t, got, 2, 3)
}

func TestPresettlementFilterBySelfPayAboveKeepsSource(t *testing.T) {
	items := samplePresettlements()
	got := NewPresettlementCompareService().FilterBySelfPayAbove(items, 50)
	assertSourceIntact(t, items)
	assertIDs(t, got, 2, 3)
}

func TestPresettlementFilterByDeductibleBelowKeepsSource(t *testing.T) {
	items := samplePresettlements()
	got := NewPresettlementCompareService().FilterByDeductibleBelow(items, 300)
	assertSourceIntact(t, items)
	assertIDs(t, got, 2)
}
