package problems

import (
    //"fmt"
    "github.com/powell0/projecteuler/utilities/math/number"
    "math/big"
    "strconv"
)

type problem0112 struct {
    id int
    description string
}

func (p problem0112) ID() int {
    return p.id
}

func (p problem0112) Description() string {
    return p.description
}

func (p problem0112) Solve() string {
    target := big.NewRat(99, 100)
    total := int64(0)
    bouncyCount := int64(0)
    notFound := true
    n := uint64(0)

    for notFound {
        n++

        if number.IsBouncyNumber(n) {
            bouncyCount++
        }

        total++

        notFound = target.Cmp(big.NewRat(bouncyCount, total)) != 0
    }

    return strconv.FormatUint(n, 10)
}

func init() {
    Registry[112] = problem0112{112, "Find the least number for which the proportion of bouncy numbers is exactly 99%."}
}
