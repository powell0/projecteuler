package problems

import (
    "github.com/powell0/projecteuler/utilities/math/factoring"
    "github.com/powell0/projecteuler/utilities/math/primes"
    "strconv"
)

type problem0077 struct {
    id int
    description string
}

func (p problem0077) ID() int {
    return p.id
}

func (p problem0077) Description() string {
    return p.description
}

func (p problem0077) Solve() string {
    const limit = 100
    const target = 5000
    result := uint64(0)

    primesList := primes.ComputePrimes(limit)

    partitions := make([]uint64, limit+1)
    partitions[0] = 0
    partitions[1] = 1

    for n := 2; n <= limit; n++ {
        sum := SumOfDistincPrimeFactors(uint64(n), primesList)
        
        for j := 1; j < n; j++ {
            distinctPrimeSum := SumOfDistincPrimeFactors(uint64(j), primesList)
            sum += (distinctPrimeSum*partitions[n-j])
        }

        sum /= uint64(n)

        if sum > target {
            result = uint64(n)
            break
        }

        partitions[n] = sum
    }
    
    return strconv.FormatUint(result, 10)
}

func SumOfDistincPrimeFactors(n uint64, primesList []uint64) uint64 {
    primeFactors := factoring.ComputePrimeFactors(n, primesList)

    sum := uint64(0)

    for _, factor := range primeFactors {
        sum += uint64(factor.Number)
    }

    return sum
}

func init() {
    Registry[77] = problem0077{77, "What is the first value which can be written as the sum of primes in over five thousand different ways?"}
}
