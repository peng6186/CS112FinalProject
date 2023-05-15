package libs

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

type EncryptedMessage struct {
	nonce      []byte
	ciphertext []byte
}

func Marshal(msg EncryptedMessage) []byte {
	return append(msg.nonce, msg.ciphertext...)
}

func Unmarshal(msgBytes []byte) EncryptedMessage {
	return EncryptedMessage{
		nonce:      msgBytes[:12],
		ciphertext: msgBytes[12:],
	}
}

func Encrypt(rawKey string, plaintext []byte) string {
	key, err := hex.DecodeString(rawKey)
	if err != nil {
		panic(err.Error())
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	encryptedMsg := EncryptedMessage{
		nonce:      nonce,
		ciphertext: ciphertext,
	}

	msgBytes := Marshal(encryptedMsg)

	return base64.RawURLEncoding.EncodeToString(msgBytes)
}

func Decrypt(rawKey string, rawEncryptedMsg string) []byte {
	msgBytes, err := base64.RawURLEncoding.DecodeString(rawEncryptedMsg)
	encryptedMsg := Unmarshal(msgBytes)
	key, _ := hex.DecodeString(rawKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, encryptedMsg.nonce, encryptedMsg.ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return plaintext
}
