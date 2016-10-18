package acm

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestSum4Zero(t *testing.T) {
	var start time.Time
	var a [4][]int64
	var i, j int64
	var N int64 = 4000

	// parse environment variable if given
	if nn, err := strconv.Atoi(os.Getenv("N")); err == nil {
		N = int64(nn)
	}

	// build data
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
	fmt.Printf("Sum4Zero1(a, %d)\t%d in %s\n", N, r1, time.Since(start))

	start = time.Now()
	r2 := Sum4Zero2(a, N)
	fmt.Printf("Sum4Zero2(a, %d)\t%d in %s\n", N, r2, time.Since(start))

	if r1 != r2 {
		t.Error("2 methods got different result!")
	}
}
