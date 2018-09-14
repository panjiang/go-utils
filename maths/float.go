package maths

import (
	"math"
	"strconv"
)

// ToFixed truncates to fix pecision
func ToFixed(num float64, precision int) float64 {
	zoom := math.Pow(10, float64(precision))
	return float64(int(num*zoom)) / zoom
}

// ToFixedString truncates to fix pecision without tailed zeroes
func ToFixedString(num float64, precision int) string {
	zoom := math.Pow(10, float64(precision))
	numFixed := float64(int(num*zoom)) / zoom
	return strconv.FormatFloat(numFixed, 'f', -1, 64)
}
