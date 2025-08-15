package tax

import (
	"testing"
)

func TestCalculateTaxByYear(t *testing.T) {
	tests := []struct {
		income float64
		year   string
		want   float64
	}{
		{96200, "2023", 21732},
	}

	for _, test := range tests {
		got, err := CalculateTaxByYear(test.income, "../tax_rates.json", test.year)
		if err != nil {
			t.Errorf("CalculateTaxByYear(%f, %s) = %v", test.income, test.year, err)
		}
		if got != test.want {
			t.Errorf("CalculateTaxByYear(%f, %s) = %f, want %f", test.income, test.year, got, test.want)
		}
	}
}

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		income float64
		want   float64
	}{
		{96200, 21732},
	}

	for _, test := range tests {
		got, err := CalculateTax(test.income, &TaxRates{
			Brackets: []TaxBracket{
				{MinIncome: 0, MaxIncome: 18200, Rate: 0},
				{MinIncome: 18200, MaxIncome: 45000, Rate: 0.19},
				{MinIncome: 45000, MaxIncome: 120000, Rate: 0.325},
				{MinIncome: 120000, MaxIncome: 180000, Rate: 0.37},
				{MinIncome: 180000, MaxIncome: 0, Rate: 0.45},
			},
		})
		if err != nil {
			t.Errorf("CalculateTax(%f) = %v", test.income, err)
		}
		if got != test.want {
			t.Errorf("CalculateTax(%f) = %f, want %f", test.income, got, test.want)
		}
	}
}
