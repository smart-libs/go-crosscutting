package decimal

//
//import (
//	"testing"
//
//	"github.com/shopspring/decimal"
//)
//
//func TestApplyRoundingMethod(t *testing.T) {
//	tests := []struct {
//		name      string
//		d         decimal.Decimal
//		method    RoundingMethod
//		precision int
//		expected  decimal.Decimal
//	}{
//		{
//			name:      "RoundingMethodFloor",
//			d:         decimal.NewFromFloat(1.234),
//			method:    RoundingMethodFloor,
//			precision: 2,
//			expected:  decimal.NewFromFloat(1.23),
//		},
//		{
//			name:      "RoundingMethodCeil",
//			d:         decimal.NewFromFloat(1.234),
//			method:    RoundingMethodCeil,
//			precision: 2,
//			expected:  decimal.NewFromFloat(1.24),
//		},
//		{
//			name:      "RoundingMethodTruncate",
//			d:         decimal.NewFromFloat(1.234),
//			method:    RoundingMethodTruncate,
//			precision: 2,
//			expected:  decimal.NewFromFloat(1.23),
//		},
//		{
//			name:      "RoundingMethodPromotePositive",
//			d:         decimal.NewFromFloat(1.234),
//			method:    RoundingMethodPromote,
//			precision: 2,
//			expected:  decimal.NewFromFloat(1.24),
//		},
//		{
//			name:      "RoundingMethodPromoteNegative",
//			d:         decimal.NewFromFloat(-1.234),
//			method:    RoundingMethodPromote,
//			precision: 2,
//			expected:  decimal.NewFromFloat(-1.24),
//		},
//		{
//			name:      "RoundingMethodHalfDown",
//			d:         decimal.NewFromFloat(1.235),
//			method:    RoundingMethodHalfDown,
//			precision: 2,
//			expected:  decimal.NewFromFloat(1.23),
//		},
//		{
//			name:      "RoundingMethodHalfUp",
//			d:         decimal.NewFromFloat(1.235),
//			method:    RoundingMethodHalfUp,
//			precision: 2,
//			expected:  decimal.NewFromFloat(1.24),
//		},
//		{
//			name:      "RoundingMethodRound",
//			d:         decimal.NewFromFloat(1.235),
//			method:    RoundingMethodRound,
//			precision: 2,
//			expected:  decimal.NewFromFloat(1.24),
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			result := ApplyRoundingMethod(tt.d, tt.method, tt.precision)
//			if !result.Equals(tt.expected) {
//				t.Errorf("ApplyRoundingMethod(%v, %v, %d) = %v; want %v", tt.d, tt.method, tt.precision, result, tt.expected)
//			}
//		})
//	}
//}
