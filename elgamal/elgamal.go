package elgamal

import (
	"math/big"
	"sort"

	"jcondeco-ciphers.com/auxfunctions"
)

/*
* n =>
* g =>
 */
func Generators(n int, g int) []int {
	l := 0
	r := []int{}
	p := 1

	nBig := big.NewInt(int64(n))
	gBig := big.NewInt(int64(g))

	for l != 1 {
		pBig := big.NewInt(int64(p))
		e := new(big.Int).Exp(nBig, pBig, nil)
		lBig := new(big.Int).Mod(e, gBig)

		l = int(lBig.Int64())
		r = append(r, l)
		p++
	}
	return r
}

func log(downN int, n int, maxVal int) int {
	gens := Generators(downN, maxVal)

	contains, index := auxfunctions.ContainsInt(n, gens)

	if contains {
		index = index + 1
		return index
	}
	return -1
}

func logTable(downN int, maxVal int) []int {
	logs := []int{}

	for i := 1; i < maxVal; i++ {
		val := log(downN, i, maxVal)
		logs = append(logs, val)
	}
	return logs
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func PrimitiveRoots(prime int) []int {
	good := []int{}
	x := []int{}

	for i := 1; i < prime; i++ {
		x = append(x, i)
	}

	for i := 1; i < prime; i++ {
		gens := Generators(i, prime)
		sort.Ints(gens)
		if slicesEqual(gens, x) {
			good = append(good, i)
		}
	}

	return good
}

func GenerateSecretKey(p int, g int, a int) int {
	pBig := big.NewInt(int64(p))
	gBig := big.NewInt(int64(g))
	aBig := big.NewInt(int64(a))

	e := new(big.Int).Exp(gBig, aBig, nil)
	mod := new(big.Int).Mod(e, pBig)
	mod2 := int(mod.Int64())
	return mod2
}
