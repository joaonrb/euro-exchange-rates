package entities

import (
	. "gotest.tools/assert"
	"testing"
)

var utilsData = []struct {
	sum    float64
	micros int64
}{
	{0.2, 200000},
	{1.2, 1200000},
	{0.244, 244000},
	{110.2003, 110200300},
	{0.2444444444444, 244444},
}

func TestSumToMicros(t *testing.T) {
	for _, d := range utilsData {
		r := SumToMicros(d.sum)
		Assert(t, r == d.micros, "SumToMicros(%f) got %d instead of the expected %d", d.sum, r, d.micros)
	}
}

func TestMicrosToSum(t *testing.T) {
	for _, d := range utilsData {
		r := MicrosToSum(d.micros)
		Assert(t, r == d.sum, "MicrosToSum(%d) got %f instead of the expected %f", d.micros, r, d.sum)
	}
}
