package problems

import (
    "math/big"
    "strconv"
)

type problem0065 struct {
    id int
    description string
}

func (p problem0065) ID() int {
    return p.id
}

func (p problem0065) Description() string {
    return p.description
}

func (p problem0065) Solve() string {
    firstTerm := big.NewRat(2,1)

    terms := make([]int64,0,99)

    term := int64(0)

    for i := 0; i < 99; i += 3 {
        term += 2
        terms = append(terms, 1)
        terms = append(terms, term)
        terms = append(terms, 1)
    }

    convergent := big.NewRat(terms[len(terms)-1],1)

    for i := len(terms) - 2; i >= 0; i-- {
        convergent.Inv(convergent)
        nextTerm := big.NewRat(terms[i], 1)
        convergent.Add(convergent, nextTerm)
    }
    convergent.Inv(convergent)
    convergent.Add(convergent, firstTerm)

    numerator := convergent.Num().String()

    sum := 0

    for i := 0; i < len(numerator); i++ {
        number, _ := strconv.Atoi(numerator[i:i+1])

        sum += number
    }

    return strconv.Itoa(sum)
}

func init() {
    Registry[65] = problem0065{65, "Find the sum of digits in the numerator of the 100th convergent of the continued fraction for e."}
}
