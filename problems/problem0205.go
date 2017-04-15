package problems

import (
    //"fmt"
    "github.com/powell0/projecteuler/utilities/math/number"
    "math/big"
)

type problem0205 struct {
    id int
    description string
}

func (p problem0205) ID() int {
    return p.id
}

func (p problem0205) Description() string {
    return p.description
}

func (p problem0205) Solve() string {

    peterTotals := make([]int64, 9*4 + 1)
    peterPermutations := number.PowInt64(4, 9)

    colinTotals := make([]int64, 6*6 + 1)
    colinPermutations := number.PowInt64(6, 6)

    generateDiceTotals(9, 4, 0, peterTotals)
    generateDiceTotals(6, 6, 0, colinTotals)

    result := big.NewRat(0, 1)

    for i := 0; i < len(peterTotals); i++ {

        peterProbability := big.NewRat(peterTotals[i], peterPermutations)

        colinCumulative := int64(0)

        for j := 0; j < i; j++ {
            colinCumulative += colinTotals[j]
        }

        colinProbability := big.NewRat(colinCumulative, colinPermutations)

        peterProbability.Mul(peterProbability, colinProbability)

        result.Add(result, peterProbability)
    }

    return result.FloatString(7)
}

func generateDiceTotals(dice int, sides int, total int, totals []int64) {
    if dice <= 0 {
        totals[total] += 1
        return
    }

    for i := 1; i <= sides; i++ {
        generateDiceTotals(dice-1, sides, total + i, totals)
    }
}

func init() {
    Registry[205] = problem0205{205, "Peter rolls 9d4 and Colin rolls 6d6 and compare totals: the highest total wins. The result is a draw if the totals are equal.  What is the probability that Peter beats Colin?"}
}
