package acm

func Gcd(m, n int64) int64 {
	if n == 0 {
		return m
	}
	return Gcd(n, m%n)
}
