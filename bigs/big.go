package bigs

import (
	"math"
	"math/big"
)

// IsOverUint64 是否超出Uint64
func IsOverUint64(x *big.Int) bool {
	max := new(big.Int).SetUint64(math.MaxUint64)
	return x.Cmp(max) > 0
}

// ToSafeUint64 封顶最大值，否则会溢出成小值
func ToSafeUint64(x *big.Int) uint64 {
	if IsOverUint64(x) {
		return math.MaxUint64
	}
	return x.Uint64()
}

// MustToUint64 转换为有效的Uint64，超出则抛异常
func MustToUint64(x *big.Int) uint64 {
	if IsOverUint64(x) {
		panic("overflow uint64 max")
	}
	return x.Uint64()
}
