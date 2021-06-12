package combinatorics

import "math/big"

/*Cnk big.Int*/
func Cnk(n, k int64) big.Int {
	if k >= 0 && n >= k {
		var nn, kkl, kkr, down, total big.Int
		nn.MulRange(1, n)
		kkl.MulRange(1, k)
		kkr.MulRange(1, n-k)
		down.Mul(&kkl, &kkr)
		down.Mul(&kkl, &kkr)
		total.Div(&nn, &down)
		return total
	}
	return *big.NewInt(0)
}

/*Cnk uint64*/
func CnkUint(n, k int64) uint64 {
	cBig := Cnk(n, k)
	return cBig.Uint64()
}

/*Cnk string*/
func CnkStr(n, k int64) string {
	cBig := Cnk(n, k)
	return cBig.String()
}
