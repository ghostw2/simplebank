package dbutils

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const alength = len(alphabet)

func RandomString(length int) string {
	builder := strings.Builder{}

	for range length {
		builder.WriteByte(alphabet[RandomInt(0, alength-1)])
	}
	return builder.String()
}
func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
