package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

const (
	keySize    = 256
	saltSize   = 32
	iterations = 1000
)

func Encrypt(plaintext, passphrase string) (string, error) {

	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	key := pbkdf2.Key([]byte(passphrase), salt, iterations, keySize/8, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	const gcmNonceSize = 12
	iv := make([]byte, gcmNonceSize)
	_, err = rand.Read(iv)
	if err != nil {
		return "", err
	}

	encryptedData := aesGCM.Seal(nil, iv, []byte(plaintext), nil)

	ciphertext := append(salt, iv...)
	ciphertext = append(ciphertext, encryptedData...)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(ciphertext, passphrase string) (string, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	const gcmNonceSize = 12

	salt := cipherBytes[:saltSize]
	iv := cipherBytes[saltSize : saltSize+gcmNonceSize]
	encryptedData := cipherBytes[saltSize+aes.BlockSize:]

	key := pbkdf2.Key([]byte(passphrase), salt, iterations, keySize/8, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesGCM.Open(nil, iv, encryptedData, nil)
	if err != nil {
		fmt.Println("Error while decrypting:", err)
		return "", err
	}

	return string(plaintext), nil
}
