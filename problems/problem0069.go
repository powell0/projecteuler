package problems

import (
    "github.com/powell0/projecteuler/utilities/math/primes"
    "strconv"
)

type problem0069 struct {
    id int
    description string
}

func (p problem0069) ID() int {
    return p.id
}

func (p problem0069) Description() string {
    return p.description
}

func (p problem0069) Solve() string {
    const limit = 1000000
    primesList := primes.ComputePrimes(100)

    numerator := uint64(1)

    for _, prime := range primesList {
        if numerator * prime <= limit {
            numerator *= prime
        } else {
            break
        }
    }

    return strconv.FormatUint(numerator, 10)
}

func init() {
    Registry[69] = problem0069{69, "Find the value of n <= 1,000,000 for which n/phi(n) is a maximum."}
}
