package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "github.com/powell0/projecteuler/utilities/math/primes"
    "strconv"
)

type problem0037 struct {
    id int
    description string
}

func (p problem0037) ID() int {
    return p.id
}

func (p problem0037) Description() string {
    return p.description
}

func (p problem0037) Solve() string {
    const limit = 1000000
    primesList := primes.ComputePrimes(limit)

    result := uint64(0)

    for _, prime := range primesList {
        isPrime := true

        if prime < 13 {
            continue
        }

        leftToRight := prime
        rightToLeft := prime

        for leftToRight > 0 && rightToLeft > 0 {
            leftToRight = number.TruncateLeftmostDigit(leftToRight)
            rightToLeft = number.TruncateRightmostDigit(rightToLeft)

            if (leftToRight > 0 && !primes.IsPrime(leftToRight)) || (rightToLeft > 0 && !primes.IsPrime(rightToLeft)) {
                isPrime = false
                break
            }
        }

        if isPrime {
            result += prime
        }
    }

    return strconv.FormatUint(result, 10)
}

func init() {
    Registry[37] = problem0037{37, "Find the sum of the only eleven primes that are both truncatable from left to right and right to left."}
}
