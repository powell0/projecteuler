package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "github.com/powell0/projecteuler/utilities/math/summation"
    "strconv"
)

type problem0085 struct {
    id int
    description string
}

func (p problem0085) ID() int {
    return p.id
}

func (p problem0085) Description() string {
    return p.description
}

func (p problem0085) Solve() string {
    const limit = 2000000
    maxGridX := 0

    maxRectanglesUpperbound := 0
    areaOfMaxUpperbound := 0
    maxRectanglesLowerbound := 0
    areaOfMaxLowerbound := 0

    // Find the maximum x value to consider
    // and compute the first upperbound solution
    x := 1

    for maxGridX == 0 {
        count := CountRectanglesInGrid(x, 1)
        
        if count > limit {
            maxGridX = x

            maxRectanglesUpperbound = count
            areaOfMaxUpperbound = x
        }

        x++
    }

    // Now compute the first lowerbound solution
    count := CountRectanglesInGrid(maxGridX - 1, 1)
    maxRectanglesLowerbound = count
    areaOfMaxLowerbound = maxGridX - 1

    // Now compute possible solutions between the bounds and refine them
    for y := 2; y <= maxGridX / 2; y++ {
        x = maxGridX

        for true {
            x--
            count := CountRectanglesInGrid(x, y)

            if count > limit {
                if count < maxRectanglesUpperbound {
                    maxRectanglesUpperbound = count
                    areaOfMaxUpperbound = x * y
                }
            } else {
                if count > maxRectanglesLowerbound {
                    maxRectanglesLowerbound = count
                    areaOfMaxLowerbound = x * y
                }
            }

            if count < maxRectanglesLowerbound {
                break
            }
        }
    }

    // Now determine which bound is closest to the target
    maxArea := areaOfMaxUpperbound

    if number.AbsInt(limit - maxRectanglesLowerbound) < number.AbsInt(limit - maxRectanglesUpperbound) {
        maxArea = areaOfMaxLowerbound
    }

    return strconv.Itoa(maxArea)
}

func CountRectanglesInGrid(x int, y int) int {
    count := int(summation.Sum1ToN(uint64(x)) * summation.Sum1ToN(uint64(y)))

    return count
}

func init() {
    Registry[85] = problem0085{85, "Although there exists no rectangular grid that contains exactly two million rectangles, find the area of the grid with the nearest solution."}
}
