package randstr

import (
	"math/rand"
	"time"
)

// Alphabet is the letters used in creating the random string.
var Alphabet = []rune(". !abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Get generates a random string sequence of given size.
func Get(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = Alphabet[rand.Intn(len(Alphabet))]
	}
	return string(b)
}
