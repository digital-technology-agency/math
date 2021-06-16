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

/*FactorialTree *big.Int*/
func FactorialTree(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	if n == 1 || n == 2 {
		return big.NewInt(n)
	}
	return factorialTreeExecute(big.NewInt(2), big.NewInt(n))
}

/*FactorialTreeString string*/
func FactorialTreeString(n int64) string {
	return FactorialTree(n).String()

}
func factorialTreeExecute(n, m *big.Int) *big.Int {
	if n.CmpAbs(m) == 1 {
		return big.NewInt(1)
	}
	if n.CmpAbs(m) == 0 {
		return n
	}
	var temp, tempR big.Int
	if temp.Sub(m, n).CmpAbs(big.NewInt(1)) == 0 {
		return tempR.Mul(n, m)
	}
	var mod, modR, modr, result big.Int
	mod.Add(n, m)
	modR.Quo(&mod, big.NewInt(2))
	resLeft := factorialTreeExecute(n, &modR)
	modr.Add(&modR, big.NewInt(1))
	resRight := factorialTreeExecute(&modr, m)
	return result.Mul(resLeft, resRight)
}
