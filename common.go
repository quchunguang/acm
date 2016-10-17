package acm

// Given m,n satisfies m>=n>=0.
func Gcd(m, n int64) int64 {
	if n == 0 {
		return m
	}
	return Gcd(n, m%n)
}
