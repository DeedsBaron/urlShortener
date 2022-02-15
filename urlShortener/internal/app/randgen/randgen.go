package randgen

import (
	"math/rand"
	"time"
)

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
)

func Generate() uint64 {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	return y1.Uint64()
}
