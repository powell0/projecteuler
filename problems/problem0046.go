package problems

import (
    "github.com/powell0/projecteuler/utilities/math/primes"
    "math"
    "strconv"
)

type problem0046 struct {
    id int
    description string
}

func (p problem0046) ID() int {
    return p.id
}

func (p problem0046) Description() string {
    return p.description
}

func (p problem0046) Solve() string {
    const limit = 100000

    primesList := primes.ComputePrimes(limit)
    numbersList := make([]bool, limit, limit)

    for i := 2; i < limit; i += 2 {
        numbersList[i] = true
    }

    for _, prime := range primesList {
        numbersList[prime] = true

        max := int(math.Sqrt(float64((limit - prime) / 2)))

        for i := 1; i <= max; i++ {
            number := int(prime) + 2 * i * i
            numbersList[number] = true
        }
    }

    result := 0

    for i := 3; i < len(numbersList); i++ {
        if !numbersList[i] {
            result = i
            break;
        }
    }

    return strconv.Itoa(result)
}

func init() {
    Registry[46] = problem0046{46, "What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?"}
}
