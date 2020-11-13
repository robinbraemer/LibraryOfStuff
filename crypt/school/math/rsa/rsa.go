package rsa

import (
	"math/big"
	"stuff/util"
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
	// phi(N) = Anzahl Teilerfremder Zahlen kleiner N
	// a hoch phi(N) = a mod N
	// N = prime1 * prime2
	// m = phi(N) = (prime1-1) * (prime2-1)

	prime1, prime2 := util.RandomPrime(bits), util.RandomPrime(bits)
	n := newInt().Mul(prime1, prime2)

	one := big.NewInt(1) // constant
	// m = (prime1-1)(prime2-1)
	m := CalcM(prime1, prime2)

	e := big.NewInt(int64(bits))
	for util.GCD(e, m).Cmp(one) > 0 { // while gcd(E, m) > 1
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

// m = (p-1)*(q-1)
func CalcM(p, q *big.Int) *big.Int {
	return big.NewInt(0).Mul(
		newInt().Sub(p, util.One),
		newInt().Sub(q, util.One),
	)
}

func newInt() *big.Int { return &big.Int{} }
