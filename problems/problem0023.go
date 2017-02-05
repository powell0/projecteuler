package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "strconv"
)

type abundantSumCandidate struct {
    Number uint64
    IsAbundantSum bool
}

type problem0023 struct {
    id int
    description string
}

func (p problem0023) ID() int {
    return p.id
}

func (p problem0023) Description() string {
    return p.description
}

func (p problem0023) Solve() string {
    const LIMIT = 28124

    abundantNumbers := make([]uint64, 0, LIMIT)

    for i := uint64(1); i < LIMIT; i++ {
        if number.IsAbundantNumber(i) {
            abundantNumbers = append(abundantNumbers, i)
        }
    }

    candidates := make([]abundantSumCandidate, LIMIT, LIMIT)

    for i := uint64(1); i < LIMIT; i++ {
        candidates[i] = abundantSumCandidate{i, false}
    }

    for i := 0; i < len(abundantNumbers); i++ {
        for j := 0; j < len(abundantNumbers); j++ {
            abundantSum := abundantNumbers[i] + abundantNumbers[j]

            if int(abundantSum) < len(candidates) {
                candidates[int(abundantSum)].IsAbundantSum = true
            }
        }
    }

    var result uint64 = 0

    for _, candidate := range candidates {
        if !candidate.IsAbundantSum {
            result += candidate.Number
        }
    }

    return strconv.FormatUint(result, 10)
}

func init() {
    Registry[23] = problem0023{23, "Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers."}
}
