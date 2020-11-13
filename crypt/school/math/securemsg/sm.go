package securemsg

import (
	"math/big"
	"stuff/message"
	"stuff/rsa"
	"stuff/util"
)

func DecryptMsg(m *big.Int, private *rsa.PrivateKey) string {
	return message.ToMessage(private.Decrypt(m).String())
}
func EncryptMsg(m string, public *rsa.PublicKey) *big.Int {
	return public.Encrypt(util.IntFromString(message.FromMessage(m)))
}
