package random

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

// RandString create uniq 10 characters string
func RandString() string {
	rand.Seed(time.Now().UnixNano())

	randBytes := make([]byte, 10)
	for i := range randBytes {
		randBytes[i] = charset[rand.Intn(len(charset))]
	}
	return string(randBytes)
}
