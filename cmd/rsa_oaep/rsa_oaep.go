package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {

	// Генерация ключей RSA (2048 бит)
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publickey := &privateKey.PublicKey

	// Шифруем сообщение открытым ключом (OAEP)
	message := []byte("секретное сообщение")
	label := []byte("") // Доп. метка, можно оставить пустой

	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publickey,
		message,
		label)
	if err != nil {
		panic(err)
	}

	// Расшифровка закрытым ключом
	plaintext, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		ciphertext,
		label)

	fmt.Println("Расшифровано: ", string(plaintext))
}
