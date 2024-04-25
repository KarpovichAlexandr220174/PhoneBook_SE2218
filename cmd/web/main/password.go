package main

import (
	"crypto/rand"
	"math/big"
)

func generateRandomPassword() string {
	const (
		upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		lowerChars   = "abcdefghijklmnopqrstuvwxyz"
		digitChars   = "0123456789"
		specialChars = "!@#$%^&*()-_=+,.?/:;{}[]`~"
		passwordLen  = 12
	)

	chars := upperChars + lowerChars + digitChars + specialChars
	var password string

	for i := 0; i < passwordLen; i++ {
		randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			panic(err)
		}
		password += string(chars[randIndex.Int64()])
	}

	return password
}
