package problems

import (
    "math/big"
    "strconv"
)

type problem0020 struct {
    id int
    description string
}

func (p problem0020) ID() int {
    return p.id
}

func (p problem0020) Description() string {
    return p.description
}

func (p problem0020) Solve() string {
    result := 0

    number := new(big.Int).MulRange(1,100)

    text := number.String()

    for i := 0; i < len(text); i++ {
        value, _ := strconv.Atoi(text[i:i+1])

        result += value
    }
    
    return strconv.Itoa(result)
}

func init() {
    Registry[20] = problem0020{20, "Find the sum of all the primes below two million."}
}
