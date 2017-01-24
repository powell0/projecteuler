package problems

import (
    "math"
    "strconv"
)

type problem0039 struct {
    id int
    description string
}

func (p problem0039) ID() int {
    return p.id
}

func (p problem0039) Description() string {
    return p.description
}

func (p problem0039) Solve() string {

    const PERIMETER_LIMIT uint64 = 1000

    type perimeterResult struct {
        perimeter uint64
        count uint64
    }

    var maxPerimeterResult perimeterResult
    resultChannel := make(chan perimeterResult)

    for perimeter := uint64(1); perimeter <= PERIMETER_LIMIT; perimeter++ {
        go func(perimeter uint64) {
            var count uint64 = 0

            for a := uint64(1); a < perimeter; a++ {
                for b := a + 1; b < perimeter; b++ {
                    c := math.Sqrt(float64(a*a) + float64(b*b))

                    if (a + b + uint64(c) == perimeter) && (c - math.Floor(c) < 0.000001) {
                        count++
                    }
                }
            }

            resultChannel <- perimeterResult{perimeter, count}
        }(perimeter)
    }

    for i := uint64(1); i <= PERIMETER_LIMIT; i++ {
        temp := <-resultChannel

        if temp.count > maxPerimeterResult.count {
            maxPerimeterResult = temp
        }
    }

    return strconv.FormatUint(maxPerimeterResult.perimeter, 10)
}

func init() {
    Registry[39] = problem0039{39, "For which value of the perimeter of a right triangle with integer sides, p <= 1000, is the number of solutions maximized?"}
}
