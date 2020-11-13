package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestSmallestDiv(t *testing.T) {
	data := []struct {
		given  int64
		expect int64
	}{
		{given: 34, expect: 17},
	}
	for i, d := range data {
		actual := SmallestDiv(big.NewInt(d.given)).Int64()
		assert.Equal(t, d.expect, actual, "test %d failed", i)
	}
}

func TestPrimeFactors(t *testing.T) {
	require.Equal(t, []int{7}, PrimeFactors(7))
	require.Equal(t, []int{2, 17}, PrimeFactors(34))

	require.Equal(t, []*big.Int{big.NewInt(7)}, BigPrimeFactors(big.NewInt(7)))
	require.Equal(t, []*big.Int{big.NewInt(2), big.NewInt(17)}, BigPrimeFactors(big.NewInt(34)))
}

func TestGGT(t *testing.T) {
	fmt.Println(ggt(7, 10))
}
