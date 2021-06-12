package combinatorics

import "math/big"

/*Factorial *big.Int*/
func Factorial(n int64) *big.Int {
	var nn big.Int
	return nn.MulRange(1, n)
}

/*FactorialInt int64*/
func FactorialInt(n int64) int64 {
	return Factorial(n).Int64()
}

/*FactorialString string*/
func FactorialString(n int64) string {
	return Factorial(n).String()
}
