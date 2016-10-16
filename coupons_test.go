package acm

import (
	"fmt"
	"testing"
)

func TestCoupons(t *testing.T) {
	x, m, y := Coupons(33)
	fmt.Println("Coupons(33)\t", x, m, "/", y)
	if x != 134 || m != 4071048809039 || y != 4375865239200 {
		t.Fail()
	}
}

func TestCoupons2(t *testing.T) {
	x, m, y := Coupons2(33)
	fmt.Println("Coupons2(33)\t", x, m, "/", y)
	if x != 134 || m != 4071048809039 || y != 4375865239200 {
		t.Fail()
	}
}
