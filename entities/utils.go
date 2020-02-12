package entities

// SumToMicros converts a float number to micros
func SumToMicros(sum float64) int64 {
	return int64(sum * micros)
}

// MicrosToSum converts micros to a float number
func MicrosToSum(mic int64) float64 {
	return float64(mic) / micros
}
