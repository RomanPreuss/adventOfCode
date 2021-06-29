package utils

import "math"

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func AbsFloat(x float64) float64 {
	if x < 0 {
		return x * -1
	}
	return x
}

func Rotate(x, y, degree float64) (float64, float64) {
	rad := degree * (math.Pi / 180)
	sin := math.Sin(rad)
	cos := math.Cos(rad)
	newX := (x * cos) - (y * sin)
	newY := (x * sin) + (y * cos)
	return math.Round(newX), math.Round(newY)
}
