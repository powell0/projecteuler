package problems

import "strconv"

type problem0002 struct {
    id int
    description string
}

func (p problem0002) ID() int {
    return p.id
}

func (p problem0002) Description() string {
    return p.description
}

func (p problem0002) Solve() string {
    var sum int

    prev := 1
    curr := 1

    for curr < 4000000 {
        if curr % 2 == 0 {
            sum += curr
        }

        prev, curr = curr, prev + curr
    }

    return strconv.Itoa(sum)
}

func init() {
    Registry[2] = problem0002{2, "By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms."}
}
