package math

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRSA(t *testing.T) {
	pk, err := rsa.GenerateKey(rand.Reader, 12)
	require.NoError(t, err)

	fmt.Println(pk.Size())
}
