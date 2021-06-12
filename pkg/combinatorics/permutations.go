package combinatorics

import "math/big"

/*Permutations big.Int*/
func Permutations(n int64) big.Int {
	return *Factorial(n)
}

/*Permutations int64*/
func PermutationsInt(n int64) int64 {
	return Factorial(n).Int64()
}

/*Permutations string*/
func PermutationStr(n int64) string {
	return Factorial(n).String()
}
