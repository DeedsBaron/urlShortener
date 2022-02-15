package encoder

import (
	"errors"
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

	return encodedBuilder.String()
}

func Decode(encoded string) (uint64, error) {
	var number uint64

	for i, symbol := range encoded {
		alphabeticPosition := strings.IndexRune(alphabet, symbol)

		if alphabeticPosition == -1 {
			return uint64(alphabeticPosition), errors.New("invalid character: " + string(symbol))
		}
		number += uint64(alphabeticPosition) * IntPow(length, uint64(i))
	}
	return number, nil
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
