package problems

import (
    "math/big"
    "strconv"
)

type problem0025 struct {
    id int
    description string
}

func (p problem0025) ID() int {
    return p.id
}

func (p problem0025) Description() string {
    return p.description
}

func (p problem0025) Solve() string {
    index := 2
    found := false

    termNMinusTwo := big.NewInt(1)
    termNMinusOne := big.NewInt(1)

    for !found {
        termN := big.NewInt(1)
        termN.Add(termNMinusTwo, termNMinusOne)
        index++

        if len(termN.String()) >= 1000 {
            found = true
        }

        termNMinusTwo = termNMinusOne
        termNMinusOne = termN
    }

    return strconv.Itoa(index)
}

func init() {
    Registry[25] = problem0025{25, "What is the index of the first term in the Fibonacci sequence to contain 1000 digits?"}
}
