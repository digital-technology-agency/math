package math

import "math/big"

func Factorial(n int64) int64 {
	var nn big.Int
	return nn.MulRange(1, n).Int64()
}
