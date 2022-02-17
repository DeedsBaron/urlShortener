package encoder

import (
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	length   = uint64(len(alphabet))
)

func Encode(number uint64) string {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(10)

	for ; number > 0; number = number / length {
		encodedBuilder.WriteByte(alphabet[(number % length)])
	}
	str := encodedBuilder.String()
	return str[:len(str)-1]
}

func IntPow(n, m uint64) uint64 {
	if m == 0 {
		return 1
	}
	result := n
	for i := uint64(2); i <= m; i++ {
		result *= n
	}
	return result
}
