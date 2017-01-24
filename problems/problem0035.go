package problems

import (
    "ian-ferguson/math/primes"
    "math"
    "strconv"
)

type problem0035 struct {
    id int
    description string
}

func (p problem0035) ID() int {
    return p.id
}

func (p problem0035) Description() string {
    return p.description
}

func (p problem0035) Solve() string {
    count := 0
    primesList := primes.ComputePrimes(1000000)

    for i := 0; i < len(primesList); i++ {
        prime := int(primesList[i])
        isCircular := true

        numRotations := int(math.Floor(math.Log10(float64(prime))));

        for j := 0; j < numRotations; j++ {
            prime = rotateNumberRight(prime, 1)

            if !primes.IsPrime(uint64(prime)) {
                isCircular = false
                break
            }
        }

        if isCircular {
            count++
        }
    }

    return strconv.Itoa(count)
}

func rotateNumberRight(number int, count int) int {
    exponent := math.Floor(math.Log10(float64(number)));

    for i := 0; i < count; i++ {
        number = (number % 10) * int(math.Pow(10, exponent)) + (number / 10)
    }

    return number
}

func init() {
    Registry[35] = problem0035{35, "How many circular primes are there below one million?"}
}
