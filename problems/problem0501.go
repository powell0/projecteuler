package problems

import (
    //"fmt"
    "github.com/powell0/projecteuler/utilities/math/factoring"
    //"github.com/powell0/projecteuler/utilities/math/factorial"
    //"github.com/powell0/projecteuler/utilities/math/number"
    "github.com/powell0/projecteuler/utilities/math/primes"
    "strconv"
)

type problem0501 struct {
    id int
    description string
}

func (p problem0501) ID() int {
    return p.id
}

func (p problem0501) Description() string {
    return p.description
}

func (p problem0501) Solve() string {
    const limit = 10000
    //count := 0

    primesList := primes.ComputePrimes(limit)

    primeFactorization := factoring.ComputePrimeFactors(1000000000001, primesList)

    for i := uint64(2); i <= limit; i++ {
        //primeFactorization := factoring.ComputePrimeFactors(i, primesList)

        //if factoring.CountFactors(primeFactorization) == 8 {
        //    count++
        //}
    }

    return strconv.Itoa(len(primeFactorization))
}

func init() {
    Registry[501] = problem0501{501, "Let f(n) be the count of numbers not exceeding n with exactly eight divisors.  Find f(1000000000000)."}
}
