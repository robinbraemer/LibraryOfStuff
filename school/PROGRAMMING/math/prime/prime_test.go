package main

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestPrime(t *testing.T) {
	s := time.Now()
	fmt.Println(countPrimes(1e6), time.Since(s))
}

func primes(max float64) (primes []float64) {
	for n := float64(0); n <= max; n++ {
		if isPrime(n) {
			primes = append(primes, n)
		}
	}
	return
}
func countPrimes(max float64) (count int) {
	for n := float64(0); n <= max; n++ {
		if isPrime(n) {
			count++
		}
	}
	return
}

func isPrime(n float64) bool {
	testUntil := math.Sqrt(n)
	for i := float64(2); i < testUntil; i++ {
		if math.Mod(n, i) == 0 {
			return false
		}
	}
	return true
}
