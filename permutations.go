package math

import "math/big"

func Permutations(n int64) big.Int {
	return *Factorial(n)
}

func PermutationsInt(n int64) int64 {
	return Factorial(n).Int64()
}

func PermutationStr(n int64) string {
	return Factorial(n).String()
}
