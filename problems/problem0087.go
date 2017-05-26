package problems

import (
    "github.com/powell0/projecteuler/utilities/math/primes"
    "math"
    "strconv"
)

type problem0087 struct {
    id int
    description string
}

func (p problem0087) ID() int {
    return p.id
}

func (p problem0087) Description() string {
    return p.description
}

func (p problem0087) Solve() string {
    const limit = 50000000
    count := 0

    primesList := primes.ComputePrimes(uint64(math.Sqrt(limit)))
    primePowerTriplets := make([]bool, limit)

    for _, secondPrime := range primesList {
        secondPower := secondPrime * secondPrime

        for _, thirdPrime := range primesList {
            thirdPower := thirdPrime * thirdPrime * thirdPrime

            if secondPower + thirdPower > limit - 1 {
                break;
            }

            for _, fourthPrime := range primesList {
                fourthPower := fourthPrime * fourthPrime * fourthPrime * fourthPrime

                if secondPower + thirdPower + fourthPower > limit - 1 {
                    break
                }

                primePowerTriplets[secondPower + thirdPower + fourthPower] = true
            }
        }
    }

    for _, triplet := range primePowerTriplets {
        if triplet {
            count++
        }
    }
    
    return strconv.Itoa(count)
}

func init() {
    Registry[87] = problem0087{87, "How many numbers below fifty million can be expressed as the sum of a prime square, prime cube, and prime fourth power?"}
}
