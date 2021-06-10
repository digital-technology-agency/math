package math

import "math/big"

func Factorial(n int64) *big.Int {
	var nn big.Int
	return nn.MulRange(1, n)
}

func FactorialInt(n int64) int64 {
	return Factorial(n).Int64()
}

func FactorialString(n int64) string {
	return Factorial(n).String()
}
