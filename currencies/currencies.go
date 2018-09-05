package currencies

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

const (
	precBits uint = 128 // 128bit精度
)

// Float64Max float64对应的最大值
var Float64Max *big.Float

// New creates a currency
func New() *big.Float {
	return new(big.Float).SetPrec(precBits)
}

// Zero returns 0 currency
func Zero() *big.Float {
	return FromFloat64(0)
}

// FromString gets a currency converted from string
func FromString(s string) *big.Float {
	d := new(big.Float).SetPrec(precBits)
	d.SetString(s)
	return d
}

// FromFloat64 gets a currency converted from float64
func FromFloat64(f float64) *big.Float {
	s := strconv.FormatFloat(f, 'f', 8, 64)
	return FromString(s)
}

// FixDecimal 固定小数点位数
func FixDecimal(b *big.Float) *big.Float {
	return FromString(ToString(b))
}

// FromInt64 gets a currency converted from int64
func FromInt64(i int64) *big.Float {
	s := strconv.FormatInt(i, 10)
	return FromString(s)
}

// FromBigInt converts bigint to currency bigfloat
func FromBigInt(i *big.Int) *big.Float {
	return FromString(i.String())
}

// ToString outputs a high precision string
func ToString(b *big.Float) string {
	s := strings.TrimRight(b.Text('f', 18), "0")
	return strings.TrimRight(s, ".")
}

// ToStringDec8 string with decimal 8
func ToStringDec8(b *big.Float) string {
	s := strings.TrimRight(b.Text('f', 8), "0")
	return strings.TrimRight(s, ".")
}

// TrimString outputs a high precision string
func TrimString(b string) string {
	s := strings.TrimRight(b, "0")
	return strings.TrimRight(s, ".")
}

// ToStringFromFloat64 convents a float64 value to currency string
func ToStringFromFloat64(f float64) string {
	return ToString(FromFloat64(f))
}

// ToFloat64 converts to float64
func ToFloat64(b *big.Float) float64 {
	return StrToFloat64(ToString(b))
}

// StrToFloat64 converts string to float64
func StrToFloat64(s string) float64 {
	sa := strings.Split(s, ".")
	if len(sa) > 1 {
		limit := 8
		if len(sa[1]) < limit {
			limit = len(sa[1])
		}

		sa[1] = sa[1][0:limit]
		s = strings.Join(sa, ".")
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return v
}

// Cmp compares two big float
func Cmp(a *big.Float, b *big.Float) int {
	x := FromString("1e+18")
	ix, _ := x.Mul(x, a).Int(&big.Int{})

	y := FromString("1e+18")
	iy, _ := y.Mul(y, b).Int(&big.Int{})

	return ix.Cmp(iy)
}

func init() {
	Float64Max = FromFloat64(math.MaxFloat64)
}
