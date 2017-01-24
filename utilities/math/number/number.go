package number

import "math"

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
