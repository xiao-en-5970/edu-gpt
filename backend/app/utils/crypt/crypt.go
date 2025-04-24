package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"golang.org/x/crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

func HashPassword(password string,salt1 string,salt2 string)([]byte, error){
	key:=DeriveKey(salt1,[]byte(salt2))
	return Encrypt([]byte(password),key)
}

func DeriveKey(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)
}

func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }

    return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    nonceSize := gcm.NonceSize()
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

    return gcm.Open(nil, nonce, ciphertext, nil)
}