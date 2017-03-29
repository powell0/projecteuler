package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "strconv"
)

type problem0074 struct {
    id int
    description string
}

func (p problem0074) ID() int {
    return p.id
}

func (p problem0074) Description() string {
    return p.description
}

func (p problem0074) Solve() string {
    const limit = 1000000
    chainLength := make(map[uint64]int)

    chainLength[0] = 2
    chainLength[1] = 1
    chainLength[2] = 1
    chainLength[145] = 1
    chainLength[169] = 3
    chainLength[871] = 2
    chainLength[872] = 2
    chainLength[1454] = 3
    chainLength[40585] = 1
    chainLength[45361] = 2
    chainLength[45362] = 2
    chainLength[363601] = 3
    
    for i := uint64(0); i < limit; i++ {
        _, ok := chainLength[i]

        if !ok {
            length := 1
            notDone := true
            n := i

            for notDone {
                digits := number.GetDigits(n)
                nextN := uint64(0)

                for digit, count := range digits {
                    digitFactorial := number.Factorial(uint64(digit))
                    digitFactorial *= uint64(count)
                    nextN += digitFactorial
                }

                nextLength, ok := chainLength[nextN]

                if ok {
                    length += nextLength
                    notDone = false
                } else {
                    length++
                    n = nextN
                }
            }

            chainLength[i] = length
        }
    }

    count := 0

    for i := uint64(0); i < limit; i++ {
        if chainLength[i] == 60 {
            count++
        }
    }

    return strconv.Itoa(count)
}

func init() {
    Registry[74] = problem0074{74, "How many chains, with a starting number below one million, contain exactly sixty non-repeating terms?"}
}
