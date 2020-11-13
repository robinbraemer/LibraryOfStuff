package modgame

import (
	"fmt"
	"math"
	"testing"
)

func TestMod(t *testing.T) {
	f(3, 0)
	fmt.Println()
	f(4, 0)
	fmt.Println()
	f(5, 0)
	fmt.Println()
	f(6, 0)
	fmt.Println()
}

func f(n float64, i int) float64 {
	if i > 20 {
		return n
	}
	fmt.Print(n, " -> ")
	i++
	if math.Mod(n, 2) == 0 {
		return f(n/2, i)
	}
	return f(n*3+1, i)
}
