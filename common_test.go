package acm

import (
	"testing"
)

var gcdTests = []struct {
	m, n int64
	out  int64
}{
	{1, 1, 1},
	{243532, 0, 243532},
	{7, 3, 1},
	{8, 4, 4},
	{8, 6, 2},
	{37279462087332, 366983722766, 564958},
	{3823485236523624, 43882834845621, 3},
}

func TestGcd(t *testing.T) {
	for _, tt := range gcdTests {
		out := Gcd(tt.m, tt.n)
		if out != tt.out {
			t.Error(tt, "->", out)
		}
	}
}
