package util

import (
	"testing"

	"github.com/blueship581/gbinsureapi/internal/constants"
)

func TestSettlementCalculator_NewRuralPolicy(t *testing.T) {
	calc := NewSettlementCalculator()
	result, err := calc.Calculate(constants.InsuranceTypeNewRural, 1000, []FeeInput{{ItemCode: "DRUG01", ItemName: "阿莫西林", MedicalCategory: constants.MedicalCategoryClassA, Amount: 1000}})
	if err != nil {
		t.Fatal(err)
	}
	if result.Deductible != 200 {
		t.Fatalf("Deductible = %v, want 200", result.Deductible)
	}
	if result.InsurancePayAmount != 520 {
		t.Fatalf("InsurancePayAmount = %v, want 520", result.InsurancePayAmount)
	}
}
