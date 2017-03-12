package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "github.com/powell0/projecteuler/utilities/math/primes"
    "strconv"
)

type problem0243 struct {
    id int
    description string
}

func (p problem0243) ID() int {
    return p.id
}

func (p problem0243) Description() string {
    return p.description
}

func (p problem0243) Solve() string {
    const limitNum = 15499
    const limitDenom = 94744
    limit := float64(limitNum)/float64(limitDenom)

    primesList := primes.ComputePrimes(100)

    denominator := uint64(1)
    startingNumber := uint64(0)
    endingNumber := uint64(0)

    for _, prime := range primesList {
        denominator *= prime
        resilience := number.Resilience(denominator)

        if resilience < limit {
            startingNumber = denominator / prime
            endingNumber = denominator
            break
        }
    }

    for n := uint64(1); startingNumber * n < endingNumber; n++ {
        denominator = startingNumber * n
        resilience := number.Resilience(denominator)
        
        if resilience < limit {
            break
        }
    }

    return strconv.FormatUint(denominator, 10)
}

func init() {
    Registry[243] = problem0243{243, "Find the smallest denominator d, having a resilience R(d) < 15499/94744"}
}
