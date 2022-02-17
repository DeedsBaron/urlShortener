package randgen

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
)

func Generate() uint64 {
	for {
		x1 := rand.NewSource(time.Now().UnixNano())
		y1 := rand.New(x1)
		res := y1.Uint64()
		if len(strconv.FormatUint(res, 10)) < 20 {
			continue
		} else {
			return res
		}
	}
}
