package problems

import (
    "github.com/powell0/projecteuler/utilities/math/primes"
    "strconv"
)

type problem0027 struct {
    id int
    description string
}

func (p problem0027) ID() int {
    return p.id
}

func (p problem0027) Description() string {
    return p.description
}

func (p problem0027) Solve() string {
    
    const LIMIT = int64(1000)
    bestA := int64(0)
    bestB := int64(0)

    maxPrimes := 0

    primesList := primes.ComputePrimes(uint64(LIMIT))

    for _, b := range primesList {
        for a := int64(-b); a < LIMIT; a++ {
            isPrime := true
            tempMaxPrimes := 0
            n := int64(0)

            for isPrime {
                result := n*n + a*n + int64(b)

                if result > 0 && primes.IsPrime(uint64(result)) {
                    tempMaxPrimes++
                    n++
                } else {
                    isPrime = false
                }
            }

            if tempMaxPrimes > maxPrimes {
                maxPrimes = tempMaxPrimes
                bestA = a
                bestB = int64(b)
            }
        }
    }

    product := bestA * bestB

    return strconv.FormatInt(product, 10)
}

func init() {
    Registry[27] = problem0027{27, "Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for consecutive values of n, starting with n=0."}
}

