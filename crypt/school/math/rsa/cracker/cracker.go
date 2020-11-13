package cracker

import (
	"fmt"
	"math/big"
	"stuff/rsa"
	"stuff/util"
)

func ComputePrivateKey(public *rsa.PublicKey) *rsa.PrivateKey {
	//p := util.SmallestDiv(public.N)
	pfs := util.BigPrimeFactors(public.N)
	if len(pfs) == 0 {
		panic("no prime factor")
	}
	fmt.Println("number prime factors:", len(pfs))
	p := pfs[len(pfs)-1] // use last prime factor
	fmt.Println("prime1:", p)

	// n = p * q
	// q = n / p
	q := new(big.Int).Div(public.N, p)
	fmt.Println("prime2:", q)

	m := rsa.CalcM(p, q)
	fmt.Println("m:", m)

	// x * d + y * m = 1 mod m
	x := new(big.Int) // will be the private key
	new(big.Int).GCD(x, new(big.Int), public.E, m)

	// It may happen that GCD gets us a negative x.
	if x.Sign() == -1 {
		// We can correct this with: x = x mod m
		x.Mod(x, m)
	}

	return &rsa.PrivateKey{
		D:         x,
		PublicKey: public,
	}
}
