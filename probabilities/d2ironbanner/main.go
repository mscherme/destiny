// Calculates the odds of recieving all Iron Banner gear from 40 packages.
// 8 weapons + 5 armor pieces
package main

import (
	"fmt"
	"math/big"
)

func H(n float64) float64 {
	var r float64
	for i := float64(1); i <= n; i++ {
		r += 1/i
	}
	return r
}

func success(n, k int64) *big.Int {
	return binomial(n-1, k-1)
}

func total(n, k int64) *big.Int {
	return binomial(n+k-1, k-1)
}

func pSuccess(n, k int64) *big.Rat {
	return big.NewRat(0, 1).SetFrac(success(n, k), total(n, k))
}

func factorial(n int64) *big.Int {
	r := big.NewInt(1)
	for i := int64(2); i <=n; i++ {
		r.Mul(r, big.NewInt(i))
	}
	return r
}

func binomial(n, k int64) *big.Int {
	b := big.NewInt(0)
	b.Binomial(n, k)
	return b
}

func ratPow(a *big.Rat, b int64) *big.Rat {
	copyA := big.NewRat(1, 1)
	for i := int64(0); i < b; i++ {
		copyA.Mul(copyA, a)
	}
	return copyA
}

func inclusionExclusionSuccessP(n, k int64) *big.Rat {
	r := big.NewRat(0, 1)
	for i := int64(1); i <= k; i++ {
		result := ratPow(big.NewRat(-1, 1), i+1)
		result.Mul(result, big.NewRat(1, 1).SetFrac(binomial(k, i), big.NewInt(1)))
		result.Mul(result, ratPow(big.NewRat(k-i,k), n))
		r.Add(r, result)
	}
	return r.Add(big.NewRat(1, 1), r.Neg(r))
}

func main() {
	var n, k int64 = 40, 13
	//fmt.Println(pSuccess(n, k).FloatString(10))
	//fmt.Println(success(n, k), total(n, k))
	fmt.Println("Expected packages to get all:", float64(k)*H(float64(k)))

	p := inclusionExclusionSuccessP(n, k)
	fmt.Println("P:", p.FloatString(10))

	totalPossibleOutcomes := ratPow(big.NewRat(k, 1), n)
	fmt.Println("Total:", totalPossibleOutcomes.Num())
	fmt.Println("Success:", big.NewRat(1, 1).Mul(totalPossibleOutcomes, p).Num())
}

