package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func IntFromString(s string) *big.Int {
	i, ok := big.NewInt(0).SetString(s, 0)
	if !ok {
		panic(fmt.Sprintf("%s to big int error", s))
	}
	return i
}

func RandomPrime(bits int) *big.Int {
	p, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		panic(fmt.Sprintf("gen prime error: %v", err))
	}
	return p
}

var One = big.NewInt(1)

// returns smallest divisor of i
func SmallestDiv(i *big.Int) (d *big.Int) {
	scrap := new(big.Int)
	zero := new(big.Int)
	add := big.NewInt(1)
	d = big.NewInt(1)

	if scrap.Mod(i, big.NewInt(2)).Cmp(zero) == 0 {
		add.Add(add, One) // +1 = 2
	}

	for {
		d.Add(d, add)
		// Is i / d even?
		if scrap.Mod(i, d).Cmp(zero) == 0 {
			return d
		}
	}
}

// Get all prime factors of a given number n
func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func BigPrimeFactors(n *big.Int) (pfs []*big.Int) {
	n = new(big.Int).Set(n) // copy n
	scratch := new(big.Int)
	zero := big.NewInt(0)
	two := big.NewInt(2)

	// Get the number of 2s that divide n
	for scratch.Mod(n, two).Cmp(zero) == 0 {
		pfs = append(pfs, two)
		n.Div(n, two)
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := big.NewInt(3); scratch.Mul(i, i).Cmp(n) <= 0; i.Add(i, two) {
		// while i divides n, append i and divide n
		for scratch.Mod(n, i).Cmp(zero) == 0 {
			pfs = append(pfs, i)
			n.Div(n, i)
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n.Cmp(two) > 0 {
		pfs = append(pfs, n)
	}

	return
}

func GCD(a, b *big.Int) *big.Int {
	return big.NewInt(0).GCD(&big.Int{}, &big.Int{}, a, b)
}

// greatest common divisor (GCD) via Euclidean algorithm
func IntGCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func ggt(a, b int) int {
	if b == 0 {
		return a
	}
	return ggt(b, a%b)
}
