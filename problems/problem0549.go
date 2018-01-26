package problems

import (
    //"fmt"
    //"github.com/powell0/projecteuler/utilities/math/factoring"
    //"github.com/powell0/projecteuler/utilities/math/factorial"
    //"github.com/powell0/projecteuler/utilities/math/number"
    //"github.com/powell0/projecteuler/utilities/math/primes"
    "strconv"
)

type problem0549 struct {
    id int
    description string
}

func (p problem0549) ID() int {
    return p.id
}

func (p problem0549) Description() string {
    return p.description
}

func (p problem0549) Solve() string {
    const limit = 100
    //primesList := primes.ComputePrimes(limit)
    //primeFactorCounts := make(map[uint64]uint64)

    for i := 2; i <= limit; i++ {

    }

    return strconv.FormatUint(0, 10)
}

func init() {
    Registry[549] = problem0549{549, "Let s(n) be the smallest number m such that n divides m!.  Let S(n) be the sum of s(i) for 2 <= i <= n.  Find S(100000000)."}
}
