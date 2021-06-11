package combinatorics

import "math/big"

func Placement(n, k int64) *big.Int {
	cnk := Cnk(n, k)
	permutations := Permutations(k)
	var down big.Int
	return down.Mul(&cnk, &permutations)
}

func PlacementInt(n, k int64) int64 {
	return Placement(n, k).Int64()
}

func PlacementStr(n, k int64) string {
	return Placement(n, k).String()
}