package acm

import (
	"testing"
)

type testcase struct {
	m, n   int64
	result int64
}

func TestGcd(t *testing.T) {
	testcases := []testcase{
		{1, 1, 1},
		{243532, 0, 243532},
		{7, 3, 1},
		{8, 4, 4},
		{8, 6, 2},
		{37279462087332, 366983722766, 564958},
		{3823485236523624, 43882834845621, 3},
	}
	for _, tc := range testcases {
		result := Gcd(tc.m, tc.n)
		if result != tc.result {
			t.Error(tc, "->", result)
		}
	}
}
