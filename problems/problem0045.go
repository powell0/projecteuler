package problems

import (
    "ian-ferguson/math/number"
    "ian-ferguson/math/summation"
    "strconv"
)

type problem0045 struct {
    id int
    description string
}

func (p problem0045) ID() int {
    return p.id
}

func (p problem0045) Description() string {
    return p.description
}

func (p problem0045) Solve() string {

    var n uint64 = 285
    var result uint64
    found := false

    for !found {
        n++

        result = summation.Sum1ToN(n)

        if number.IsPentagonaNumber(result) && number.IsHexagonalNumber(result) {
            found = true
        }
    }

    return strconv.FormatUint(result, 10)
}

func init() {
    Registry[45] = problem0045{45, "Find the next triangle number that is also pentagonal and hexagonal."}
}
