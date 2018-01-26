package problems

import (
    "fmt"
    "github.com/powell0/projecteuler/utilities/math/factoring"
    "github.com/powell0/projecteuler/utilities/math/number"
    "github.com/powell0/projecteuler/utilities/math/primes"
    "math"
    "math/big"
    //"strconv"
)

type problem0590 struct {
    id int
    description string
}

func (p problem0590) ID() int {
    return p.id
}

func (p problem0590) Description() string {
    return p.description
}

func (p problem0590) Solve() string {
    const limit = 50000

    result := leastCommonMultipleofRange(1, limit)

    return result.String()
}

func leastCommonMultipleofRange (start uint64, end uint64) big.Int {
    primeLimit := uint64(math.Sqrt(float64(end)) + 1)
    primesList := primes.ComputePrimes(primeLimit)

    factors := make(map[uint64]uint64)

    for i := start; i <= end; i++ {
        tempFactors := factoring.ComputePrimeFactors(uint64(i), primesList)

        for _, factor := range tempFactors {
            index := factor.Number
            factors[index] = number.MaxUint64(factors[index], uint64(factor.Count))
        }
    }

    result := big.NewInt(1)

    for factor, count := range factors {
        product := big.NewInt(int64(factor))
        product.Exp(product, big.NewInt(int64(count)), nil)

        result.Mul(result, product)
    }

    fmt.Println(len(factors))

    return *result
}

func init() {
    // Registry[590] = problem0590{590, "Let H(n) denote the number of sets of positive integers such that the least common multiple of the integers in the set equals n. Let L(n) denote the least common multiple of the numbers 1 through n. Let HL(n) denote H(L(n)).  Find HL(50000). Give your answer modulo 10^9."}
}
