package decimal

type RoundingMethod int

const (
	// RoundingMethodNever - Don't do any rounding, simply truncate and print a warning in case of a remainder.
	// Otherwise the same as RoundTrunc
	RoundingMethodNever = RoundingMethod(0)

	// RoundingMethodFloor - Round to the largest integral value not greater than given decimal.
	// e.g. 0.5 -> 0.0 and -0.5 -> -1.0
	RoundingMethodFloor = RoundingMethod(1)

	// RoundingMethodCeil - Round to the smallest integral value not less than @p this.
	// e.g. 0.5 -> 1.0 and -0.5 -> -0.0
	RoundingMethodCeil = RoundingMethod(2)

	// RoundingMethodTruncate - no rounding, simply truncate any fraction
	RoundingMethodTruncate = RoundingMethod(3)

	// RoundingMethodPromote - Use RoundCeil for positive and RoundFloor for negative values of @p this.
	// e.g. 0.5 -> 1.0 and -0.5 -> -1.0
	RoundingMethodPromote = RoundingMethod(4)

	// RoundingMethodHalfDown - Round up or down with the following constraints:
	// 0.1 .. 0.5 -> 0.0 and 0.6 .. 0.9 -> 1.0
	RoundingMethodHalfDown = RoundingMethod(5)

	// RoundingMethodHalfUp - Round up or down with the following constraints:
	// 0.1 .. 0.4 -> 0.0 and 0.5 .. 0.9 -> 1.0
	RoundingMethodHalfUp = RoundingMethod(6)

	// RoundingMethodRound - Use RoundHalfDown for 0.1 .. 0.4 and RoundHalfUp for 0.6 .. 0.9.
	// Use RoundHalfUp for 0.5 in case the resulting numerator is odd, RoundHalfDown in case the resulting
	// numerator is even.
	// e.g. 0.5 -> 0.0 and 1.5 -> 2.0
	RoundingMethodRound = RoundingMethod(7)
)

//func ApplyDecimalConfig(config Options, d types.Decimal) types.Decimal {
//	return ApplyRoundingMethod(d, config.RoundingMethod, config.Precision)
//}
//
//func ApplyRoundingMethod(d types.Decimal, method RoundingMethod, precision int) types.Decimal {
//	switch method {
//	//case RoundingMethodFloor:
//	//	return d.RoundFloor(int32(precision))
//	//case RoundingMethodCeil:
//	//	return d.RoundCeil(int32(precision))
//	//case RoundingMethodNever, RoundingMethodTruncate:
//	//	return d.Truncate(int32(precision))
//	//case RoundingMethodPromote:
//	//	if d.IsPositive() {
//	//		return ApplyRoundingMethod(d, RoundingMethodCeil, precision)
//	//	}
//	//	return ApplyRoundingMethod(d, RoundingMethodFloor, precision)
//	//case RoundingMethodHalfDown:
//	//	return d.RoundDown(int32(precision))
//	//case RoundingMethodHalfUp:
//	//	return d.RoundUp(int32(precision))
//	//case RoundingMethodRound:
//	//	return d.Round(int32(precision))
//	default:
//		return d
//	}
//}
