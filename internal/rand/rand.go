package rand

import (
	"math"
	"math/rand"
)

const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const NumberBytes = "1234567890"

func GenerateString(alphabet string, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[rand.Int63()%int64(len(alphabet))]
	}
	return string(b)
}
func GenerateNumber(n int) int {
	min := int(math.Pow10(n))
	max := int(math.Pow10(n + 1))
	return min + rand.Intn(max-min)
}
