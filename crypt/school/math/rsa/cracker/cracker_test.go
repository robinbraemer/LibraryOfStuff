package cracker

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"stuff/rsa"
	"stuff/rsa/testdata"
	"stuff/securemsg"
	"stuff/util"
	"testing"
	"time"
)

func TestCrack(t *testing.T) {
	private := rsa.GenerateKey(7)
	fmt.Println("actual D:", private.D)

	cracked := ComputePrivateKey(private.PublicKey)
	fmt.Println("cracked D:", cracked.D)

	require.Equal(t, private.D, cracked.D)
}

func TestCrackEasy(t *testing.T) {
	private := rsa.GenerateKey(15)
	fmt.Println("Actual private:",
		"D:", private.D,
		"N:", private.N,
		"E:", private.E)

	cracked := ComputePrivateKey(private.PublicKey)
	require.Equal(t, private.D, cracked.D)

	msg := securemsg.EncryptMsg("BAAAM", private.PublicKey)
	fmt.Println("msg:", msg)
	crackedMsg := securemsg.DecryptMsg(msg, cracked)
	fmt.Println(crackedMsg)
}

func TestCrackedMessage(t *testing.T) {
	start := time.Now()
	cracked := ComputePrivateKey(testdata.Maik)
	fmt.Println(time.Since(start), "cracked D:", cracked.D)
	msg := util.IntFromString("154235978395484938785329")
	securemsg.DecryptMsg(msg, cracked)
}
