package bits

func CountSetBits(n uint64) int {
    count := 0

    for count = 0; n > 0; count++ {
        n &= (n-1)
    }

    return count
}

