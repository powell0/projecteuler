package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "strconv"
)

type problem0044 struct {
    id int
    description string
}

func (p problem0044) ID() int {
    return p.id
}

func (p problem0044) Description() string {
    return p.description
}

func (p problem0044) Solve() string {
    const limit = 10001
    result := uint64(1000000000)

    pentagonalNumbers := make([]uint64, limit, limit)

    for i := 1; i < limit; i++ {
        pentagonalNumbers[i] = uint64(i*(3*i-1)/2)
    }

    for i := limit - 1; i > 0; i-- {
        for j := i - 1; j > 0; j-- {
            if number.IsPentagonaNumber(pentagonalNumbers[i] + pentagonalNumbers[j]) && 
               number.IsPentagonaNumber(pentagonalNumbers[i] - pentagonalNumbers[j]) {
                   result = number.MinUint64(result, pentagonalNumbers[i] - pentagonalNumbers[j])
            }
        }
    }

    return strconv.FormatUint(result, 10)
}

func init() {
    Registry[44] = problem0044{44, "Find the pair of pentagonal numbers, Pj and Pk, for which their sum and difference are pentagonal and D = |Pk - Pj| is minimised; what is the value of D?"}
}
