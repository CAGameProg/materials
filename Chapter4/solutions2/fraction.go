package main

import (
	"fmt"
	"strconv"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

type Fraction struct {
	num   int
	denom int
}

func NewFraction(num, denom int) *Fraction {
	s := new(Fraction)
	s.num = num
	s.denom = denom

	s.simplify()

	return s
}

func (f *Fraction) simplify() {
	gcd := gcd(f.num, f.denom)
	f.num /= gcd
	f.denom /= gcd
}

func (a *Fraction) Mul(b *Fraction) *Fraction {
	return NewFraction(a.num*b.num, a.denom*b.denom)
}

func (a *Fraction) Div(b *Fraction) *Fraction {
	return NewFraction(a.num*b.denom, a.denom*b.num)
}

func (a *Fraction) Add(b *Fraction) *Fraction {
	return NewFraction(a.num*b.denom+b.num*a.denom, a.denom*b.denom)
}

func (a *Fraction) Sub(b *Fraction) *Fraction {
	return NewFraction(a.num*b.denom-b.num*a.denom, a.denom*b.denom)
}

func (f *Fraction) String() string {
	return strconv.Itoa(f.num) + "/" + strconv.Itoa(f.denom)
}

func (f *Fraction) Approx() float64 {
	return float64(f.num) / float64(f.denom)
}

func main() {
	frac1 := NewFraction(16, 14)
	frac2 := NewFraction(33, 12)

	result := frac1.Mul(frac2)
	fmt.Println(frac1, "*", frac2, "=", result, "≈", result.Approx())
}
