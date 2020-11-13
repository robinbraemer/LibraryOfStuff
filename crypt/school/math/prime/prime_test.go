package prime

import (
	"fmt"
	"testing"
	"time"
)

func TestPrime(t *testing.T) {
	s := time.Now()
	fmt.Println(CountPrimes(0, 1e6), time.Since(s))
}
