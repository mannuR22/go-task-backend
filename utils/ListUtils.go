package utils

import (
	"crypto/rand"
	"math/big"
)

func GetRefId() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < 16; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result += string(charset[randomIndex.Int64()])
	}
	return result
}

func IsStringInArray(arr []string, target string) bool {
	for _, str := range arr {
		if str == target {
			return true
		}
	}

	return false
}
