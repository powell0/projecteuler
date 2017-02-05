package problems

import (
    "github.com/powell0/projecteuler/utilities/math/factoring"
    "strconv"
)

type problem0021 struct {
    id int
    description string
}

func (p problem0021) ID() int {
    return p.id
}

func (p problem0021) Description() string {
    return p.description
}

func (p problem0021) Solve() string {
    const LIMIT = 10000
    sum := uint64(0)

    sumOfFactors := make([]uint64, LIMIT, LIMIT)

    sumOfFactors[1] = 1

    for i := uint64(2); i < LIMIT; i++ {
        factors := factoring.ComputeFactors(i)

        for _, factor := range factors {
            if factor != i {
                sumOfFactors[i] += factor
            }
        }
    }

    for i := uint64(2); i < LIMIT; i++ {
        if sumOfFactors[i] < LIMIT && sumOfFactors[sumOfFactors[i]] == i && i != sumOfFactors[i] {
            sum += i
        }
    }

    return strconv.FormatUint(sum, 10)
}

func init() {
    Registry[21] = problem0021{21, "Evaluate the sum of all the amicable numbers under 10000."}
}
