package pkg

import (
	"backupAgent/domain/pkg/log"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
)

func Encryption(filePath string) (string, error) {
	plaintext, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	key := []byte("1233214567893332")
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	// Save back to file
	cyName := GetFilePath(filePath) + ".bak"
	err = ioutil.WriteFile(cyName, ciphertext, 0777)
	if err != nil {
		return "", err
	}
	log.Logger.Infof("加密成功,加密文件%s", cyName)
	return cyName, nil
}
