package acm

// Gcd needs m,n satisfies m>=n>=0 and returns gcd(m,n).
func Gcd(m, n int64) int64 {
	if n == 0 {
		return m
	}
	return Gcd(n, m%n)
}
