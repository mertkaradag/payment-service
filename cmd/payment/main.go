package main

import (
	"github.com/k0kubun/pp/v3"
	"payment-service/internal/bank"
	"payment-service/internal/strategy"
)

func main() {
	akbank := &bank.Akbank{
		Commission: 1.5,
	}
	halkbank := &bank.Halkbank{
		Commission: 1.5,
	}
	isbank := &bank.Isbank{
		Commission: 1.06,
	}

	banks := []bank.Bank{akbank, halkbank, isbank}

	lowestCommissionStrategy := &strategy.LowestCommissionStrategy{}
	stg := strategy.New(lowestCommissionStrategy)

	lowestCommissionBank := stg.Execute(banks)

	for _, b := range banks {
		pp.Printf("Bank: %s, Commission: %v\n", b.FetchName(), b.FetchCommission())
	}
	pp.Printf("Lowest Commission Bank: %s\n", lowestCommissionBank.FetchName())

}
