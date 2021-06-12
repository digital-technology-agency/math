package combinatorics

import "math/big"

/*Permutations big.Int*/
func Placement(n, k int64) *big.Int {
	cnk := Cnk(n, k)
	permutations := Permutations(k)
	var down big.Int
	return down.Mul(&cnk, &permutations)
}

/*PlacementInt int64*/
func PlacementInt(n, k int64) int64 {
	return Placement(n, k).Int64()
}

/*PlacementStr string*/
func PlacementStr(n, k int64) string {
	return Placement(n, k).String()
}
