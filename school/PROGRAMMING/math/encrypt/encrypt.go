package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type PrivateKey struct {
	D *big.Int
	*PublicKey
}

type PublicKey struct {
	N, E *big.Int
}

func (pub *PublicKey) Encrypt(message *big.Int) *big.Int {
	x := newInt().Exp(message, pub.E, nil)
	return x.Mod(x, pub.N)
}

func (pri *PrivateKey) Decrypt(x *big.Int) *big.Int {
	// x^privateKey mod N
	return newInt().Exp(x, pri.D, pri.N)
}

func GenerateKey(bits int) *PrivateKey {
	// gruppentheorie
	// a hoch prime1 = a mod prime1
	// a hoch prime1-1 = 1 mod prime1
	//
	// phi -> eulersche Phi-Funktion
	// phi(N) = Anzahl Teilerfremnder Zahlen kleiner N
	// a hoch phi(N) = a mod N
	// N = prime1 * prime2
	// m = phi(N) = (prime1-1) * (prime2-1)

	prime1, prime2 := prime(bits), prime(bits)
	n := newInt().Mul(prime1, prime2)

	one := big.NewInt(1) // constant
	// m = (prime1-1)(prime2-1)
	m := big.NewInt(0).Mul(
		newInt().Sub(prime1, one),
		newInt().Sub(prime2, one),
	)

	e := big.NewInt(int64(bits))
	for gcd(e, m).Cmp(one) > 0 { // while gcd(E, m) > 1
		e.Add(e, one)
	}
	d := newInt().ModInverse(e, m)

	return &PrivateKey{
		D: d,
		PublicKey: &PublicKey{
			N: n,
			E: e,
		},
	}
}

func newInt() *big.Int { return &big.Int{} }

func prime(bits int) *big.Int {
	p, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		panic(fmt.Sprintf("gen prime error: %v", err))
	}
	return p
}

func gcd(a, b *big.Int) *big.Int {
	return big.NewInt(0).GCD(&big.Int{}, &big.Int{}, a, b)
}

func xOfGcd(a, b *big.Int) (x *big.Int) {
	x = big.NewInt(0)
	big.NewInt(0).GCD(x, &big.Int{}, a, b)
	return
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd2(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
