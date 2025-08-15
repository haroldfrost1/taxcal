package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/haroldfrost1/taxcal/tax"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Get tax year
	fmt.Print("Please enter the income year: ")
	scanner.Scan()
	taxYear := strings.TrimSpace(scanner.Text())
	if taxYear == "" {
		fmt.Println("Error: Tax year cannot be empty")
		os.Exit(1)
	}

	// Validate tax year format (YYYY-YYYY)
	if !strings.Contains(taxYear, "-") || len(taxYear) != 9 {
		fmt.Println("Error: Tax year must be in format YYYY-YYYY")
		os.Exit(1)
	}

	years := strings.Split(taxYear, "-")
	startYear, err1 := strconv.Atoi(years[0])
	endYear, err2 := strconv.Atoi(years[1])

	if err1 != nil || err2 != nil || len(years[0]) != 4 || len(years[1]) != 4 {
		fmt.Println("Error: Tax year must be in format YYYY-YYYY with valid years")
		os.Exit(1)
	}

	if endYear != startYear+1 {
		fmt.Println("Error: End year must be the year after start year")
		os.Exit(1)
	}

	// Get income
	fmt.Print("Please enter your total taxable income for the full income year: ")
	scanner.Scan()
	incomeStr := strings.TrimSpace(scanner.Text())

	income, err := strconv.ParseFloat(incomeStr, 64)
	if err != nil {
		fmt.Printf("Error: Invalid income amount '%s'\n", incomeStr)
		os.Exit(1)
	}

	if income < 0 {
		fmt.Println("Error: Income cannot be negative")
		os.Exit(1)
	}

	// Convert end year to string for tax calculation
	endYearStr := strconv.Itoa(endYear)

	taxAmount, err := tax.CalculateTaxByYear(income, "tax_rates.json", endYearStr)
	if err != nil {
		fmt.Printf("Error calculating tax: %v\n", err)
		os.Exit(1)
	}

	afterTax := income - taxAmount
	effectiveRate := (taxAmount / income) * 100

	fmt.Printf("\n--- Tax Calculation for %s ---\n", taxYear)
	fmt.Printf("Taxable Income: $%.2f\n", income)
	fmt.Printf("Tax Amount:     $%.2f\n", taxAmount)
	fmt.Printf("After Tax:      $%.2f\n", afterTax)
	fmt.Printf("Effective Rate: %.2f%%\n", effectiveRate)
}
