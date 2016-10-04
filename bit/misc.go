package bit

import (
	"math"
)

func Length(n uint) uint {
	return uint(math.Floor(math.Log2(float64(n)) + 1))
}
