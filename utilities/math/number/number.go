package number

import (
    "github.com/powell0/projecteuler/utilities/math/factoring"
    "math"
)

func MinInt(x, y int) int {
    if x < y {
        return x
    }

    return y
}

func MinInt64(x, y int64) int64 {
    if x < y {
        return x
    }

    return y
}

func MinUint64(x, y uint64) uint64 {
    if x < y {
        return x
    }

    return y
}

func MaxInt(x, y int) int {
    if x > y {
        return x
    }

    return y
}

func MaxInt64(x, y int64) int64 {
    if x > y {
        return x
    }

    return y
}

func MaxUint64(x, y uint64) uint64 {
    if x > y {
        return x
    }

    return y
}

func PowInt(x, y int) int {
    result := 1

    for i := 0; i < y; i++ {
        result *= x
    }

    return result
}

func PowUint64(x, y uint64) uint64 {
    result := uint64(1)

    for i := uint64(0); i < y; i++ {
        result *= x
    }

    return result
}

func IsTriangleNumber(number uint64) bool {
    root := (-1 + math.Sqrt(8 * float64(number) + 1)) / 2

    return (math.Abs(math.Floor(root + 0.5) - root) < 0.000001)
}

func IsPentagonaNumber(number uint64) bool {
    root := (1 + math.Sqrt(24 * float64(number) + 1)) / 6

    return (math.Abs(math.Floor(root + 0.5) - root) < 0.000001)
}

func IsHexagonalNumber(number uint64) bool {
    root := (1 + math.Sqrt(8 * float64(number) + 1)) / 4

    return (math.Abs(math.Floor(root + 0.5) - root) < 0.000001)
}

func IsAbundantNumber(number uint64) bool {
    factors := factoring.ComputeFactors(number)

    sum := uint64(0)

    for _, factor := range factors {
        if factor != number {
            sum += factor
        }
    }

    return sum > number
}

func TruncateLeftmostDigit(number uint64) uint64 {
    var result uint64

    if number < 10 {
        result = 0
    } else {
        base := uint64(math.Pow(10, math.Ceil(math.Log10(float64(number))) - 1))

        result = number % base
    }

    return result
}

func TruncateRightmostDigit(number uint64) uint64 {
    return number / 10
}

func GetDigits(number uint64) map[int]int {
    results := make(map[int]int)

    for number > 0 {
        digit := int(number % 10)
        number /= 10

        results[digit]++
    }

    return results
}

func IsPermutation(first map[int]int, second map[int]int) bool {
    result := true

    for k, v := range first {
        if second[k] != v {
            result = false
            break
        }

        delete(second, k)
    }

    if len(second) != 0 {
        result = false
    }

    return result
}

func HasUniqueDigits(number uint64) bool {
    result := true

    digits := make([]bool, 10)

    for number > 0 {
        digit := int(number % 10)
        number /= 10

        if digits[digit] {
            result = false

            break;
        }

        digits[digit] = true
    }

    return result
}

func NumbersAreUnique(a, b uint64) bool {
    result := true

    aDigits := GetDigits(a)
    bDigits := GetDigits(b)

    for k, _ := range aDigits {
        _, ok := bDigits[k]

        if ok {
            result = false
            break
        }
    }

    return result
}

func CountDigits(number uint64) int {
    return int(math.Floor(math.Log10(float64(number))) + 1)
}
