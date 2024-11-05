package strategy

import "payment-service/internal/bank"

// CommissionStrategy is the interface that all commission strategies should implement
type CommissionStrategy interface {
	Execute(banks []bank.Bank) bank.Bank
}

// Strategy is the struct that holds the commission strategy
type Strategy struct {
	commissionStrategy CommissionStrategy
}

// New creates a new Strategy
func New(commissionStrategy CommissionStrategy) *Strategy {
	return &Strategy{commissionStrategy: commissionStrategy}
}

// Execute executes the commission strategy
func (c *Strategy) Execute(banks []bank.Bank) bank.Bank {
	return c.commissionStrategy.Execute(banks)
}
