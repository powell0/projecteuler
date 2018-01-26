package problems

import (
    //"fmt"
    //"github.com/powell0/projecteuler/utilities/math/number/pandigital"
    "strconv"
)

type problem0571 struct {
    id int
    description string
}

func (p problem0571) ID() int {
    return p.id
}

func (p problem0571) Description() string {
    return p.description
}

func (p problem0571) Solve() string {
    
    return strconv.FormatUint(0, 10)
}

func init() {
    //Registry[571] = problem0571{571, ""}
}
