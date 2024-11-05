package strategy_test

import (
	"payment-service/internal/bank"
	"payment-service/internal/strategy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowestCommissionStrategy_Execute(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		banks    []bank.Bank
		expected *float64
	}{
		{
			name: "single bank",
			banks: []bank.Bank{
				&bank.Akbank{Commission: 5.0},
			},
			expected: newFloat64(5.0),
		},
		{
			name: "multiple banks",
			banks: []bank.Bank{
				&bank.Akbank{Commission: 3.0},
				&bank.Isbank{Commission: 7.0},
				&bank.Halkbank{Commission: 5.0},
			},
			expected: newFloat64(3.0),
		},
		{
			name:     "no banks",
			banks:    []bank.Bank{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			stg := &strategy.LowestCommissionStrategy{}
			result := stg.Execute(tt.banks)

			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, result.FetchCommission())
			}
		})
	}
}
