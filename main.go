package main

import (
	"fmt"
	"math/big"

	"github.com/jeremyletang/bigintbenches/bn256"
	"github.com/ncw/gmp"
)

func main() {
	var a, b bn256.Element
	a.SetUint64(42)
	b.SetUint64(7)
	fmt.Printf("%v\n", a.Mul(&a, &b))
}

func FibGoff(n int) *bn256.Element {

	var left, right bn256.Element
	left.SetUint64(1)
	right.SetUint64(1)

	var helper1, helper2 bn256.Element
	helper1.SetUint64(0)
	helper2.SetUint64(0)

	var two bn256.Element
	two.SetUint64(2)

	bin := fmt.Sprintf("%b", n)
	for i := 1; i < len(bin); i++ {

		helper1.Mul(&two, &right)
		helper1.Sub(&helper1, &left)
		helper1.Mul(&left, &helper1)

		helper2.Mul(&right, &right)
		left.Mul(&left, &left)
		helper2.Add(&helper2, &left)

		if bin[i] == '0' {
			left.Set(&helper1)
			right.Set(&helper2)
		} else {
			left.Set(&helper2)
			right.Add(&helper1, &helper2)
		}
	}

	return &left
}

func FibBigInt(n int) *big.Int {

	left := big.NewInt(1)
	right := big.NewInt(1)

	helper1 := big.NewInt(0)
	helper2 := big.NewInt(0)

	bin := fmt.Sprintf("%b", n)
	for i := 1; i < len(bin); i++ {

		helper1.Mul(big.NewInt(2), right)
		helper1.Sub(helper1, left)
		helper1.Mul(left, helper1)

		helper2.Mul(right, right)
		left.Mul(left, left)
		helper2.Add(helper2, left)

		if bin[i] == '0' {
			left.Set(helper1)
			right.Set(helper2)
		} else {
			left.Set(helper2)
			right.Add(helper1, helper2)
		}
	}

	return left
}

func FibGmpInt(n int) *gmp.Int {

	left := gmp.NewInt(1)
	right := gmp.NewInt(1)

	helper1 := gmp.NewInt(0)
	helper2 := gmp.NewInt(0)

	bin := fmt.Sprintf("%b", n)
	for i := 1; i < len(bin); i++ {

		helper1.Mul(gmp.NewInt(2), right)
		helper1.Sub(helper1, left)
		helper1.Mul(left, helper1)

		helper2.Mul(right, right)
		left.Mul(left, left)
		helper2.Add(helper2, left)

		if bin[i] == '0' {
			left.Set(helper1)
			right.Set(helper2)
		} else {
			left.Set(helper2)
			right.Add(helper1, helper2)
		}
	}

	return left
}

func FibInt(n int) int64 {

	var left int64 = 1
	var right int64 = 1

	var helper1 int64 = 0
	var helper2 int64 = 0

	bin := fmt.Sprintf("%b", n)
	for i := 1; i < len(bin); i++ {

		helper1 = 2 * right
		helper1 = helper1 - left
		helper1 = left * helper1

		helper2 = right * right
		left = left * left
		helper2 = helper2 + left

		if bin[i] == '0' {
			left = helper1
			right = helper2
		} else {
			left = helper2
			right = helper1 + helper2
		}
	}

	return left
}
