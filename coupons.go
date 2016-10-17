/*
Coupons
UVa 10288

The solution equals to calculate Sum(n/i), i in [1..n].
*/
package acm

import (
	"math/big"
)

// Coupons using int64, should care overflow carefully. N=33 is OK.
func Coupons(N int64) (int64, int64, int64) {
	var i int64
	var x, y int64 = 0, 1
	for i = 1; i <= N; i++ {
		x = i*x + y*N
		y = i * y
		gcd := Gcd(x, y)
		x /= gcd
		y /= gcd
		// fmt.Println(i, ":", x, y)
	}

	return x / y, x % y, y
}

// Coupons2 using math/big library, will never overflow.
func Coupons2(N int64) (int64, int64, int64) {
	var i int64
	n := big.NewRat(0, 1)
	for i = 1; i <= N; i++ {
		n.Add(n, big.NewRat(N, i))
		// fmt.Println(i, ":", n.Num(), n.Denom())
	}
	x, y := n.Num(), n.Denom()
	m := big.NewInt(0)
	x.DivMod(x, y, m)
	return x.Int64(), m.Int64(), y.Int64()
}
