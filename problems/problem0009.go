package problems

import (
    "math"
    "strconv"
)

type problem0009 struct {
    id int
    description string
}

func (p problem0009) ID() int {
    return p.id
}

func (p problem0009) Description() string {
    return p.description
}

func (p problem0009) Solve() string {

    var a uint64
    var b uint64
    var c float64
    var product uint64

    const sum uint64 = 1000

    found := false

    for a = 1; a < sum && !found; a++ {
        for b = a + 1; b < sum && !found; b++ {
            c = math.Sqrt(float64(a*a) + float64(b*b))

            if (a + b + uint64(c) == sum) && (c - math.Floor(c) < 0.000001) {
                product = a * b * uint64(c)
                found = true
            }
        }
    }

    return strconv.FormatUint(product, 10)
}

func init() {
    Registry[9] = problem0009{9, "There exists exactly one Pythagorean triplet for which a + b + c = 1000.  Find the product abc."}
}
