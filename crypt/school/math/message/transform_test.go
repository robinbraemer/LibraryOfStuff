package message

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTrans(t *testing.T) {
	require.Equal(t,
		"Hallo",
		ToMessage(FromMessage("Hallo")))
}
