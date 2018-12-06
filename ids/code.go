package ids

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strings"

	"github.com/tv42/base58"
)

const zoomRange int64 = 10000000

// UniqueRandCode 唯一随机码
// 算法: I放大R倍数后随机偏移[0,R)
func UniqueRandCode(i int64) string {
	m := UniqueRandID(i)
	rb := base58.EncodeBig(nil, m)
	return string(rb)
}

// GenerateCDKey 生成指定数量CDkey
func GenerateCDKey(n int) {
	mp := map[int64]string{}

	i := 1
	for {
		ri := rand.Int63n(74141784152535)

		_, ok := mp[ri]
		if ok {
			continue
		}

		rb := base58.EncodeBig(nil, big.NewInt(ri))
		rs := string(rb)
		fmt.Printf("%s\n", rs)
		mp[ri] = rs
		i++
		if len(mp) == n {
			break
		}
	}
}

// FixLenRandCDKey 定长纯随机CDKEY，需要手动判断重复
func FixLenRandCDKey(length int) (string, error) {
	if length > 10 {
		return "", errors.New("max length 10")
	}
	rangeTop := math.Pow(58, float64(length))
	r := rand.Int63n(int64(rangeTop))
	s := string(base58.EncodeBig(nil, big.NewInt(r)))
	if len(s) == length {
		return s, nil
	}
	paddingLeft := strings.Repeat("1", length-len(s))
	return paddingLeft + s, nil
}
