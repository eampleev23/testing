package main

import (
	"crypto/rand"
	"crypto/subtle"
	"fmt"
	"golang.org/x/crypto/argon2"
)

func main() {
	// 1. Пользователь вводит пароль
	password := "пароль-рыба-меч"

	// 2. Генерируем случайную соль
	salt := make([]byte, 16)
	rand.Read(salt)

	// 3. Производим ключ длиной 32 байта (AES-256)
	masterKey := argon2.IDKey([]byte(password), salt, 3, 64*1024, 8, 32)

	// 4. Проверка ключа (иммитация входа)
	enteredPassword := "пароль-рыба-меч"
	testKey := argon2.IDKey([]byte(enteredPassword), salt, 3, 64*1024, 8, 32)
	if subtle.ConstantTimeCompare(testKey, masterKey) == 1 {
		fmt.Println("Пароль верный! Ключ:", masterKey)
	} else {
		fmt.Println("Неверный пароль")
	}
}
