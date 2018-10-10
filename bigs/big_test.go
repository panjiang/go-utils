package bigs

import (
	"math"
	"math/big"
	"testing"
)

func TestToSafeUint64(t *testing.T) {
	// 最大-1
	b := new(big.Int).SetUint64(math.MaxUint64 - 1)
	t.Log(IsOverUint64(b))
	t.Log(ToSafeUint64(b))
	t.Log(MustToUint64(b))

	// 达到最大
	b.Add(b, new(big.Int).SetUint64(1))
	t.Log(IsOverUint64(b))
	t.Log(ToSafeUint64(b))
	t.Log(MustToUint64(b))

	// 溢出
	b.Add(b, new(big.Int).SetUint64(1))
	t.Log(IsOverUint64(b))
	t.Log(ToSafeUint64(b))
	t.Log(MustToUint64(b))
}
