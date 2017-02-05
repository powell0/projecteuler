package problems

import (
    "github.com/powell0/projecteuler/utilities/math/factoring"
    "github.com/powell0/projecteuler/utilities/math/primes"
    "github.com/powell0/projecteuler/utilities/math/summation"
    "strconv"
)

type problem12 struct {
    id int
    description string
}

func (p problem12) ID() int {
    return p.id
}

func (p problem12) Description() string {
    return p.description
}

func (p problem12) Solve() string {
    const PRIME_LIMIT = 1000
    const FACTOR_LIMIT = 500
    result := uint64(0)
    primeNumbers := primes.ComputePrimes(PRIME_LIMIT)
    n := uint64(1)

    for result == 0 {
        n++

        var factors1 []factoring.PrimeFactor
        var factors2 []factoring.PrimeFactor

        if n % 2 == 0 {
            factors1 = factoring.ComputePrimeFactors(n / 2, primeNumbers)
            factors2 = factoring.ComputePrimeFactors(n + 1, primeNumbers)
        } else {
            factors1 = factoring.ComputePrimeFactors(n, primeNumbers)
            factors2 = factoring.ComputePrimeFactors((n + 1) / 2, primeNumbers)
        }

        if (factoring.CountFactors(factors1) * factoring.CountFactors(factors2) > FACTOR_LIMIT) {
            result = summation.Sum1ToN(n);
        }
    }

    return strconv.FormatUint(result, 10)
}

func init() {
    Registry[12] = problem12{12, "What is the value of the first triangle number to have over five hundred divisors?"}
}
