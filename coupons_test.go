package acm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoupons(t *testing.T) {
	x, m, y := Coupons(33)
	fmt.Println("Coupons(33)\t", x, m, "/", y)

	assert.EqualValues(t, 134, x)
	assert.EqualValues(t, 4071048809039, m)
	assert.EqualValues(t, 4375865239200, y)
}

func TestCoupons2(t *testing.T) {
	x, m, y := Coupons2(33)
	fmt.Println("Coupons2(33)\t", x, m, "/", y)

	assert.EqualValues(t, 134, x)
	assert.EqualValues(t, 4071048809039, m)
	assert.EqualValues(t, 4375865239200, y)
}
