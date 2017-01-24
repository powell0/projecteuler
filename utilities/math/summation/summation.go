package summation

func Sum1ToN(n uint64) uint64 {
    return n * (n + 1) / 2
}

func SumMultiplesOfN(n uint64, limit uint64) uint64 {
    return Sum1ToN((limit - 1) / n) * n
}

func SumOfSquares(n uint64) uint64 {
    return n * (n + 1) * (2 * n + 1) / 6
}
