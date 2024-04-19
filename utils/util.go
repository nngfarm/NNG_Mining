package utils

import (
	"github.com/shopspring/decimal"
	"math/big"
)

// PowDecimal Pow Decimal
func PowDecimal(decimals int32) decimal.Decimal {
	return decimal.NewFromInt(10).Pow(decimal.NewFromInt32(decimals))
}

// WeiToAmount wei to decimals
func WeiToAmount(iValue interface{}, decimals uint8) decimal.Decimal {
	// new Big Int
	value := new(big.Int)
	switch v := iValue.(type) {
	case string:
		value.SetString(v, 0)
	case *big.Int:
		value = v
	case decimal.Decimal:
		value = v.BigInt()
	}
	num := decimal.NewFromBigInt(value, 0)
	// To wei
	result := num.DivRound(PowDecimal(int32(decimals)), 18)

	return result
}

// AmountToWei decimals to wei
func AmountToWei(iAmount interface{}, decimals uint8) *big.Int {
	return AmountToWeiDecimal(iAmount, decimals).BigInt()
}

func AmountToWeiDecimal(iAmount interface{}, decimals uint8) decimal.Decimal {
	// amount zero
	amount := decimal.NewFromInt32(0)
	// IAmount
	switch v := iAmount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int:
		amount = decimal.NewFromInt(int64(v))
	case int64:
		amount = decimal.NewFromInt(v)
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}
	return amount.Mul(PowDecimal(int32(decimals)))
}
