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
<<<<<<< HEAD
	return y1.Uint64()
=======
	return uint64(y1.Intn(9223372036854775807))
>>>>>>> ed8f4a1 (postgresql container is configured and working)
}
