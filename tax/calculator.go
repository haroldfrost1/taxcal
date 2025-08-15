package tax

import (
	"encoding/json"
	"fmt"
	"os"
)

// TaxBracket represents a tax bracket with rate and threshold
type TaxBracket struct {
	MinIncome float64 `json:"min_income"`
	MaxIncome float64 `json:"max_income"`
	Rate      float64 `json:"rate"`
}

// TaxRates represents the complete tax rate structure
type TaxRates struct {
	Brackets []TaxBracket `json:"brackets"`
}

// YearlyTaxRates represents tax rates organized by year
type YearlyTaxRates map[string]TaxRates

// CalculateTax calculates tax based on income and rates
func CalculateTax(income float64, rates *TaxRates) (float64, error) {
	if rates == nil {
		return 0, fmt.Errorf("tax rates cannot be nil")
	}

	var totalTax float64
	for _, bracket := range rates.Brackets {
		if income <= bracket.MinIncome {
			continue
		}

		taxableInThisBracket := income
		if bracket.MaxIncome > 0 && income > bracket.MaxIncome {
			taxableInThisBracket = bracket.MaxIncome
		}
		taxableInThisBracket -= bracket.MinIncome

		totalTax += taxableInThisBracket * bracket.Rate
	}

	return totalTax, nil
}

// CalculateTaxByYear calculates tax based on income, rates file, and end year
// This reads the rates file, unmarshals it, and then calculates the tax with the rates for the given year
func CalculateTaxByYear(income float64, ratesFile string, endYear string) (float64, error) {
	rates, err := loadTaxRatesByYear(ratesFile, endYear)
	if err != nil {
		return 0, fmt.Errorf("failed to load tax rates: %w", err)
	}
	return CalculateTax(income, rates)
}

// loadTaxRatesByYear loads tax rates for a specific year from a JSON file
func loadTaxRatesByYear(filename string, endYear string) (*TaxRates, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var yearlyRates YearlyTaxRates
	if err := json.Unmarshal(data, &yearlyRates); err != nil {
		return nil, err
	}

	rates, exists := yearlyRates[endYear]
	if !exists {
		return nil, fmt.Errorf("tax rates not found for year %s", endYear)
	}

	return &rates, nil
}
