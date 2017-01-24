package problems

import (
    "strconv"
)

type problem0014 struct {
    id int
    description string
}

func (p problem0014) ID() int {
    return p.id
}

func (p problem0014) Description() string {
    return p.description
}

func (p problem0014) Solve() string {
    maxCount := uint64(0)
    bestN := uint64(0)

    counts := make(map[uint64]uint64)
    counts[1] = 1

    for n := uint64(2); n < 1000000; n++ {
        currentNumber := n
        currentCount := uint64(0)
        nextCount := uint64(0)
        found := false

        for !found {
            currentCount++
            currentNumber = nextCollatzNumber(currentNumber)

            nextCount, found = counts[currentNumber]
        }

        counts[n] = currentCount + nextCount
    }

    for n := uint64(1); n < 1000000; n++ {
        if counts[n] > maxCount {
            maxCount = counts[n]
            bestN = n
        }
    }
    
    return strconv.FormatUint(bestN, 10)
}

func init() {
    Registry[14] = problem0014{14, "Which starting number, under one million, produces the longest Collatz sequence?"}
}

func nextCollatzNumber(n uint64) uint64 {
    result := uint64(0)

    if n % 2 == 0 {
        result = n / 2
    } else {
        result = 3 * n + 1
    }

    return result
}
