package problems

import (
    "math/big"
    "strconv"
)

type problem0053 struct {
    id int
    description string
}

func (p problem0053) ID() int {
    return p.id
}

func (p problem0053) Description() string {
    return p.description
}

func (p problem0053) Solve() string {
    const limit = 100

    count := 0

    threshold := big.NewInt(1000000)

    for n := int64(1); n <= limit; n++ {
        for k := int64(1); k <= n; k++ {
            z := new (big.Int).Binomial(n, k)

            if z.Cmp(threshold) > 0 {
                count++
            }
        }
    }

    return strconv.Itoa(count)
}

func init() {
    Registry[53] = problem0053{53, "How many, not necessarily distinct, values of nCr, for 1 <= n <= 100, are greater than one-million?"}
}
