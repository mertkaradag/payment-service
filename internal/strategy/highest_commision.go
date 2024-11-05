package strategy

import (
	"payment-service/internal/bank"
)

// HighestCommissionStrategy is the strategy that selects the bank with the highest commission
type HighestCommissionStrategy struct{}

// Execute selects the bank with the highest commission
func (l *HighestCommissionStrategy) Execute(banks []bank.Bank) bank.Bank {
	var highestCommissionBank bank.Bank
	highestCommission := float64(0)

	for _, b := range banks {
		commission := b.FetchCommission()
		if commission > highestCommission {
			highestCommission = commission
			highestCommissionBank = b
		}
	}

	return highestCommissionBank
}
