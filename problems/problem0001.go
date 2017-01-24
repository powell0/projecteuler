package problems

import "strconv"

type problem0001 struct {
    id int
    description string
}

func (p problem0001) ID() int {
    return p.id
}

func (p problem0001) Description() string {
    return p.description
}

func (p problem0001) Solve() string {
    var sum int

    for i := 1; i < 1000; i++ {
        if i % 3 == 0 {
            sum += i
        } else if i % 5 == 0 {
            sum += i
        }
    }

    return strconv.Itoa(sum)
}

func init() {
    Registry[1] = problem0001{1, "Find the sum of all the multiples of 3 or 5 below 1000."}
}
