package strategy

import (
	"math"
	"payment-service/internal/bank"
)

// LowestCommissionStrategy is the strategy that selects the bank with the lowest commission
type LowestCommissionStrategy struct{}

// Execute selects the bank with the lowest commission
func (l *LowestCommissionStrategy) Execute(banks []bank.Bank) bank.Bank {
	var lowestCommissionBank bank.Bank
	lowestCommission := math.MaxFloat32

	for _, b := range banks {
		commission := b.FetchCommission()
		if commission < lowestCommission {
			lowestCommission = commission
			lowestCommissionBank = b
		}
	}

	return lowestCommissionBank
}
