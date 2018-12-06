package ids

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFixLenRandCDKey(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		str, err := FixLenRandCDKey(10)
		if err != nil {
			panic(err)
		}
		fmt.Println(str)
	}
}
