package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEncDec(t *testing.T) {
	m := newMsg("Hallo", robin.PublicKey)
	require.Equal(t, getMsg(m, robin), "Hallo")
}
