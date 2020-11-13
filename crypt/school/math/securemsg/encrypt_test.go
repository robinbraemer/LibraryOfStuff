package securemsg

import (
	"github.com/stretchr/testify/require"
	"stuff/rsa/testdata"
	"testing"
)

func TestEncDec(t *testing.T) {
	m := EncryptMsg("Hallo", testdata.Robin.PublicKey)
	require.Equal(t, DecryptMsg(m, testdata.Robin), "Hallo")
}
