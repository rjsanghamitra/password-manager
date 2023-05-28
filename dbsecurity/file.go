package dbsecurity

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	mrand "math/rand"
	"pwdmgr/error"
)

var iv []byte

func Encrypt(data string, key string) (encrypted string) {
	iv = []byte("1234567890123456")
	block, err := aes.NewCipher([]byte(key))
	error.CheckError(err)

	// cfb is a mode of operating in cryptography
	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(data))
	cfb.XORKeyStream(cipherText, []byte(data))
	return base64.StdEncoding.EncodeToString(cipherText)
}

func Decrypt(encrypted string, key string) string {
	iv = []byte("1234567890123456")
	block, err := aes.NewCipher([]byte(key))
	error.CheckError(err)

	cipherText, err := base64.StdEncoding.DecodeString(encrypted)
	error.CheckError(err)

	cfb := cipher.NewCFBEncrypter(block, iv)
	decryptedBytes := make([]byte, len(cipherText))
	cfb.XORKeyStream(decryptedBytes, cipherText)
	return string(decryptedBytes)
}

func RandomPasswordGenerator() (pwd string) {
	// the character set from which the characters that will be there in the password will be selected.
	charset := []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890@#$%&*")
	b := make([]byte, 14)
	for i := range b {
		b[i] = charset[mrand.Intn(len(charset))]
	}
	return string(b)
}
