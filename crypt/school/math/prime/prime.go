package prime

import (
	"math"
)

func ForPrimes(min, max int, fn func(p int) bool) {
	for i := min; i <= max; i++ {
		if IsPrime(i) && !fn(i) {
			return
		}
	}
}
func IsPrime(n int) bool {
	testUntil := int(math.Sqrt(float64(n)))
	for i := 2; i < testUntil; i++ {
		if math.Mod(float64(n), float64(i)) == 0 {
			return false
		}
	}
	return true
}

func Primes(min, max int) (primes []int) {
	ForPrimes(min, max, func(p int) bool {
		primes = append(primes, p)
		return true
	})
	return
}
func CountPrimes(min, max int) (count int) {
	ForPrimes(min, max, func(p int) bool {
		count++
		return true
	})
	return
}
