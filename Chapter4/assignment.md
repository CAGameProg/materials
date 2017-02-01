# Pointers and Structs

## 1. vector.go

Write a struct definition for a vector which is simply two numbers x and y. These two components
should both be `float64`.

Write the following methods for it:

* `func (v *Vector) Magnitude() float64`: computes the magnitude of the vector.
* `func (v *Vector) Angle() float64`: returns the angle (in degrees) of the vector with the x-axis.
* `func (v *Vector) Mul(scalar float64) *Vector`: multiplies a given scalar with the vector.
* `func (a *Vector) Add(b *Vector) *Vector`: adds a vector to the vector.
* `func (a *Vector) Dot(b *Vector) float64`: computes the dot product between two vectors.
* `func (a *Vector) AngleBetween(b *Vector) float64`: computes the angle between two vectors. Hint: use the dot product.

## 2. fraction.go

Write a `Fraction` struct which can be used to store irrational numbers. A fraction has two components,
a numerator and a denominator, which should both be integers.

Using this struct, write the following methods for it:

* `func (a *Fraction) Mul(b *Fraction) *Fraction`: returns the fraction which is the result of multiplying a and b together.
* `func (a *Fraction) Div(b *Fraction) *Fraction`: returns the fraction which is the result of dividing a and b.
* `func (a *Fraction) Add(b *Fraction) *Fraction`: returns the fraction which is the result of adding a and b together.
* `func (a *Fraction) Sub(b *Fraction) *Fraction`: returns the fraction which is the result of subtracting b from a.
* `func (f *Fraction) String() string`: returns the a string representation of the fraction, for example: `2/3`.
* `func (f *Fraction) Approx() float64`: returns the fraction's decimal approximation.

Also write a constructor:

* `func NewFraction(num, denom int) *Fraction`

All fractions should be fully simplified. To help with this, you probably want to write
a private method: `func (f *Fraction) simplify()`. 

Hint: To simplify a fraction, you need to divide the numerator and the denominator by the
fraction's greatest common divisor. You can write the function `gcd(a, b int) int` to find the
greatest common divisor between two numbers using [Euclid's algorithm](https://en.wikipedia.org/wiki/Greatest_common_divisor#Using_Euclid.27s_algorithm).