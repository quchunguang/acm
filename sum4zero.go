package acm

import (
	"sort"
)

/*
Sum4Zero - 4 values whose sum is 0
SWERC 2005
Level 2
4 values from 4 N-size array (N<=4000), whose sum is 0. How many choices?
*/
func Sum4Zero(a [4][]int64, N int64) (ret int64) { return Sum4Zero1(a, N) }

// Sum4Zero1 performance based on len(counter), which based on the range of values of a.
func Sum4Zero1(a [4][]int64, N int64) (ret int64) {
	var i, j int64
	counter := make(map[int64]int64)

	for i = 0; i < N; i++ {
		for j = 0; j < N; j++ {
			// add combine of first two array, N*N items at most
			counter[a[0][i]+a[1][j]]++
		}
	}

	for i = 0; i < N; i++ {
		for j = 0; j < N; j++ {
			// add combine of last two array, search its inverse
			// and then sum to ret
			ret += counter[-a[2][i]-a[3][j]]
		}
	}
	return
}

// Sum4Zero2 performance based on len(counter), which based on the range of values of a.
func Sum4Zero2(a [4][]int64, N int64) (ret int64) {
	var i, j int64
	// var sums []int64
	sums := make([]int64, N*N, N*N) // -0.5s than above

	for i = 0; i < N; i++ {
		for j = 0; j < N; j++ {
			// add combine of first two array, N*N items at most
			// sums = append(sums, a[0][i]+a[1][j]) // with init: make([]int64, 0, N*N)
			sums[N*i+j] = a[0][i] + a[1][j] // -0.3s than above
		}
	}

	sort.Sort(Int64Slice(sums)) // 5s

	for i = 0; i < N; i++ { // 9s
		for j = 0; j < N; j++ {
			// add combine of last two array, search its inverse
			// and then sum to ret
			key := -a[2][i] - a[3][j]
			ret += timesInInt64s(sums, key)
		}
	}
	return
}

// Int64Slice for sort
type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func timesInInt64s(a []int64, key int64) (ret int64) {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, len(a)
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
		if a[h] < key {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.

	for j := i; j < len(a) && a[j] == key; j++ {
		ret++
	}
	for j := i - 1; j >= 0 && a[j] == key; j-- {
		ret++
	}
	return
}
