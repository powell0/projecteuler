package problems

import (
    "ian-ferguson/math/primes"
    "strconv"
)

type problem0007 struct {
    id int
    description string
}

func (p problem0007) ID() int {
    return p.id
}

func (p problem0007) Description() string {
    return p.description
}

func (p problem0007) Solve() string {

    const limit = 10001
    count := 0
    var number uint64 = 1

    for count < limit {
        number++

        if primes.IsPrime(number) {
            count++
        }
    }

    return strconv.FormatUint(number, 10)
}

func init() {
    Registry[7] = problem0007{7, "What is the 10001st prime number?"}
}
