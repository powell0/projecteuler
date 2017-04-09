package problems

import (
    "math/big"
    "strconv"
)

type problem0057 struct {
    id int
    description string
}

func (p problem0057) ID() int {
    return p.id
}

func (p problem0057) Description() string {
    return p.description
}

func (p problem0057) Solve() string {
    const limit = 1000
    count := 0
    expansions := make([]*big.Rat, limit + 1)

    zero := big.NewRat(0,1)
    one := big.NewRat(1, 1)
    two := big.NewRat(2, 1)

    expansions[0] = zero
    expansions[1] = new(big.Rat).Inv(two)

    for i := 2; i <= limit; i++ {
        current := new(big.Rat).Add(two, expansions[i-1])
        expansions[i] = current.Inv(current)
    }

    for i := 0; i <= limit; i++ {
        expansions[i].Add(expansions[i], one)

        if len(expansions[i].Num().String()) > len(expansions[i].Denom().String()) {
            count++
        }
    }
    
    return strconv.Itoa(count)
}


func init() {
    Registry[57] = problem0057{57, "In the first one-thousand continued fraction expansions of the square root of two, how many fractions contain a numerator with more digits than denominator?"}
}
