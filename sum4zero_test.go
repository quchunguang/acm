package acm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

// TestSum4Zero1 runs about 10s in my laptop.
func TestSum4Zero(t *testing.T) {
	var start time.Time
	var elapsed time.Duration

	// build data
	var a [4][]int64
	var i, j int64
	const N = 4000

	rand.Seed(time.Now().UTC().UnixNano())
	for i = 0; i < 4; i++ {
		a[i] = make([]int64, N, N)
		for j = 0; j < N; j++ {
			// the range of values of a is full int32 here
			a[i][j] = rand.Int63n(1<<32) - 1<<31
		}
	}

	start = time.Now()
	r1 := Sum4Zero1(a, N)
	elapsed = time.Since(start)
	fmt.Printf("Sum4Zero1(a, %d)\t%d in %v\n", N, r1, elapsed)

	start = time.Now()
	r2 := Sum4Zero2(a, N)
	elapsed = time.Since(start)
	fmt.Printf("Sum4Zero2(a, %d)\t%d in %v\n", N, r2, elapsed)

	assert.Equal(t, r1, r2, "One of methods is wrong!")
}
