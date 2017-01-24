package problems

import (
    "ian-ferguson/math/primes"
    "strconv"
)

type problem0010 struct {
    id int
    description string
}

func (p problem0010) ID() int {
    return p.id
}

func (p problem0010) Description() string {
    return p.description
}

func (p problem0010) Solve() string {

    var result uint64 = 0

    primesList := primes.ComputePrimes(2000000)

    for i := 0; i < len(primesList); i++ {
        result += primesList[i]
    }

    return strconv.FormatUint(result, 10)
}

func init() {
    Registry[10] = problem0010{10, "Find the sum of all the primes below two million."}
}
