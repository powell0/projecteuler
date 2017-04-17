package problems

import (
    //"github.com/powell0/projecteuler/utilities/math/factoring"
    //"github.com/powell0/projecteuler/utilities/math/factorial"
    "github.com/powell0/projecteuler/utilities/math/number"
    "github.com/powell0/projecteuler/utilities/math/primes"
    "strconv"
)

type problem0214 struct {
    id int
    description string
}

func (p problem0214) ID() int {
    return p.id
}

func (p problem0214) Description() string {
    return p.description
}

func (p problem0214) Solve() string {
    const limit = 40000000
    sum := uint64(0)

    primesList := primes.ComputePrimes(limit)
    chainLengths := make(map[uint64]int)
    chainLengths[1] = 1
   

    for _, prime := range primesList {
        n := prime - 1

        chainLengths[prime] = GetChainLength(chainLengths, n) + 1

        if chainLengths[prime] == 25 {
            sum += prime
        }
    }

    return strconv.FormatUint(sum, 10)
}

func GetChainLength(chainLengths map[uint64]int, n uint64) int {
    length, ok := chainLengths[n]

    if !ok {
        nextN := number.EulerTotient(n)

        length = GetChainLength(chainLengths, nextN) + 1
        chainLengths[n] = length
    }

    return length
}

func init() {
    Registry[214] = problem0214{214, "What is the sum of all primes less than 40000000 which generate a chain of length 25?"}
}
