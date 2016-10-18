package acm

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestCoupons(t *testing.T) {
	var start time.Time
	var N int64 = 33

	// parse environment variable if given
	if nn, err := strconv.Atoi(os.Getenv("N")); err == nil {
		N = int64(nn)
	}

	start = time.Now()
	x1, m1, y1 := Coupons1(N)
	fmt.Printf("Coupons1(%d)\t%d%d / %d in %s\n", N, x1, m1, y1, time.Since(start))

	start = time.Now()
	x2, m2, y2 := Coupons2(N)
	fmt.Printf("Coupons2(%d)\t%d%d / %d in %s\n", N, x2, m2, y2, time.Since(start))

	if !(x1 == x2 && m1 == m2 && y1 == y2) {
		t.Error("2 methods got different result!")
	}
}
