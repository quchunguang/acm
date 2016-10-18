package acm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCoupons(t *testing.T) {
	var x, m, y int64
	var start time.Time
	var elapsed time.Duration

	start = time.Now()
	x, m, y = Coupons1(33)
	elapsed = time.Since(start)
	fmt.Printf("Coupons1(33)\t%d%d / %d in %v\n", x, m, y, elapsed)
	assert.EqualValues(t, 134, x)
	assert.EqualValues(t, 4071048809039, m)
	assert.EqualValues(t, 4375865239200, y)

	start = time.Now()
	x, m, y = Coupons2(33)
	elapsed = time.Since(start)
	fmt.Printf("Coupons2(33)\t%d%d / %d in %v\n", x, m, y, elapsed)
	assert.EqualValues(t, 134, x)
	assert.EqualValues(t, 4071048809039, m)
	assert.EqualValues(t, 4375865239200, y)
}
