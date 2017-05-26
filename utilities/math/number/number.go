package number

import (
    "github.com/powell0/projecteuler/utilities/math/factoring"
    "math"
)

func AbsInt(x int) int {
    if x < 0 {
        x *= -1
    }

    return x
}

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

func PowInt64(x, y int64) int64 {
    result := int64(1)

    for i := int64(0); i < y; i++ {
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

func Reverse(number uint64) uint64 {
    result := uint64(0)

    for number > 0 {
        digit := number % 10
        number /= 10

        result *= 10
        result += digit
    }

    return result
}

func AreCoprime(a uint64, b uint64) bool {
    coprime := true

    if a == b {
        coprime = false
    } else if a == 1 || b == 1 {
        coprime = true
    } else {
        commonFactors := make(map[uint64]struct{})

        primeFactorizationA := factoring.ComputePrimeFactors(a, nil)
        primeFactorizationB := factoring.ComputePrimeFactors(b, nil)

        for _, primeFactor := range primeFactorizationA {
            commonFactors[primeFactor.Number] = struct{}{}
        }

        for _, primeFactor := range primeFactorizationB {
            _, ok := commonFactors[primeFactor.Number]

            if ok {
                coprime = false
                break
            }
        }
    }

    return coprime
}

func EulerTotient(n uint64) uint64 {
    primeFactorization := factoring.ComputePrimeFactors(n, nil)

    count := float64(n)

    for _, primeFactor := range primeFactorization {
        count *= (1 - 1 / float64(primeFactor.Number))
    }

    return uint64(count)
}

func Resilience(n uint64) float64 {
    count := EulerTotient(n)

    resilience := float64(count) / float64(n-1)

    return resilience
}

func GCD(a uint64, b uint64) uint64 {
    if a == b {
        return a
    }

    if a == 0 {
        return 0
    }

    if b == 0 {
        return 0
    }

    if ^a & 1 == 1 {
        if b & 1 == 1 {
            return GCD(a >> 1, b)
        } else {
            return GCD(a >> 1, b >> 1) << 1
        }
    }

    if ^b & 1 == 1 {
        return GCD(a, b >> 1)
    }

    if a > b {
        return GCD((a - b) >> 1, b)
    }

    return GCD((b - a) >> 1, a)
}

func LogBase(x float64, b float64) float64 {
    return math.Log(x) / math.Log(b)
}

func Factorial(n uint64) uint64 {
    result := uint64(1)

    for i := n; i > 0; i-- {
        result *= i
    }

    return result
}
