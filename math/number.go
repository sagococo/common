package math

import "math"

func IsPrime(value int) bool {
	if value < 2 {
		return false
	}
	end := int(math.Sqrt(float64(value)))
	for i := 2; i <= end; i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}
