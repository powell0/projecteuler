package problems

import (
    "strconv"
)

type problem0048 struct {
    id int
    description string
}

func (p problem0048) ID() int {
    return p.id
}

func (p problem0048) Description() string {
    return p.description
}

func (p problem0048) Solve() string {

    const MODULUS uint64 = 10000000000
    const LIMIT uint64 = 1000
    var result uint64 = 0

    for i := uint64(1); i <= LIMIT; i++ {
        temp := uint64(1)
        for j := uint64(1); j <= i; j++ {
            temp *= i
            temp %= MODULUS 
        }

        result += temp
        result %= MODULUS
    }

    return strconv.FormatUint(result, 10)
}

func init() {
    Registry[48] = problem0048{48, "Find the last ten digits of the series, 1^(1) + 2^(2) + 3^(3) + ... + 1000^(1000)."}
}
