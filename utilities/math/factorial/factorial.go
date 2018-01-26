package factorial

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "math"
)

func CalculateTrailingZeroes(value int64) int64 {
    return CalculateMultiplicityOfPrime(value, 5)
}

func CalculateMultiplicityOfPrime(value int64, prime int64) int64 {
    floatingPointValue := float64(value)
    floatingPointPrime := float64(prime)
    k := int64(math.Floor(number.LogBase(floatingPointValue, floatingPointPrime)))
    count := int64(0)

    for i := int64(1); i <= k; i++ {
        denominator := math.Pow(floatingPointPrime, float64(i))

        count += int64(math.Floor(floatingPointValue / denominator))
    }

    return count
}
