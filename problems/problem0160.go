package problems

import (
    //"fmt"
    "strconv"
)

type problem0160 struct {
    id int
    description string
}

func (p problem0160) ID() int {
    return p.id
}

func (p problem0160) Description() string {
    return p.description
}

func (p problem0160) Solve() string {
    const limit = 1000000000000

    factorial := uint64(1);
    count := uint64(0)

    const tempLimit = 1000000000

    for i := uint64(2); i < tempLimit; i++ {
        product := i

        for product % 10 == 0 {
            product /= 10
        }

        factorial *= product

        for factorial % 10 == 0 {
            factorial /= 10
        }

        factorial %= limit

        count++
    }

    return strconv.FormatUint(factorial % limit, 10)
}

func init() {
    Registry[160] = problem0160{160, "For any N, let f(N) be the last five digits before the trailing zeroes in N!.  Find f(1,000,000,000,000)."}
}
