package problems

import (
    "github.com/powell0/projecteuler/utilities/math/primes"
    "strconv"
)

type problem0187 struct {
    id int
    description string
}

func (p problem0187) ID() int {
    return p.id
}

func (p problem0187) Description() string {
    return p.description
}

func (p problem0187) Solve() string {
    const limit = 100000000
    count := 0

    primesList := primes.ComputePrimes(limit/2)

    for i := 0; i < len(primesList); i++ {
        firstPrime := primesList[i]

        for j := i; j < len(primesList); j++ {
            secondPrime := primesList[j]

            if firstPrime * secondPrime < limit {
                count++
            } else {
                break
            }
        }

    }

    return strconv.Itoa(count)
}

func init() {
    Registry[187] = problem0187{187, "How many composite integers, n < 10^8, have precisely two, not necessarily distinct, prime factors?"}
}
