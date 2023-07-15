package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func GeneratorRandNo(digit int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	var n int32
	for {
		n = r.Int31n(int32(math.Pow10(digit)))
		if n >= int32(math.Pow10(digit-1)) {
			break
		}

	}
	return fmt.Sprintf("%04v", n)
}
