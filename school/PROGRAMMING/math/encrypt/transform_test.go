package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTrans(t *testing.T) {
	require.Equal(t,
		"Hallo",
		toMessage(fromMessage("Hallo")))
}
