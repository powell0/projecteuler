package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "strconv"
)

type problem0092 struct {
    id int
    description string
}

func (p problem0092) ID() int {
    return p.id
}

func (p problem0092) Description() string {
    return p.description
}

func (p problem0092) Solve() string {

    const limit = 10000000

    count := 0

    arrivesAt1 := make(map[uint64]struct{})
    arrivesAt1[1] = struct{}{}

    arrivesAt89 := make(map[uint64]struct{})
    arrivesAt89[89] = struct{}{}

    for n := uint64(1); n < limit; n++ {
        finished := false
        candidate := n

        for !finished {
            digits := number.GetDigits(candidate)

            candidate = 0

            for digit, count := range digits {
                candidate += uint64(digit*digit*count)
            }

            _,  ok := arrivesAt1[candidate]

            if ok {
                arrivesAt1[candidate] = struct{}{}
                finished = true
                break
            }

            _,  ok = arrivesAt89[candidate]

            if ok {
                arrivesAt89[candidate] = struct{}{}
                finished = true
                count++
                break
            }
        }
    }

    return strconv.Itoa(count)
}

func init() {
    Registry[92] = problem0092{92, "How many square digit chain starting numbers below ten million will arrive at 89"}
}
