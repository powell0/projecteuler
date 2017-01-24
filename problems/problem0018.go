package problems

import (
    "github.com/powell0/projecteuler/utilities"
    "github.com/powell0/projecteuler/utilities/math/number"
    "strconv"
)

type problem0018 struct {
    id int
    description string
}

func (p problem0018) ID() int {
    return p.id
}

func (p problem0018) Description() string {
    return p.description
}

func (p problem0018) Solve() string {

    var result int64 = 0
    levels := 15
    var triangle utilities.NumberTriangle
    triangle.Initialize(levels)

    triangle.Data[0] = []int64{ 75 }
    triangle.Data[1] = []int64{ 95, 64 }
    triangle.Data[2] = []int64{ 17, 47, 82 }
    triangle.Data[3] = []int64{ 18, 35, 87, 10 }
    triangle.Data[4] = []int64{ 20, 04, 82, 47, 65 }
    triangle.Data[5] = []int64{ 19, 01, 23, 75, 03, 34 }
    triangle.Data[6] = []int64{ 88, 02, 77, 73, 07, 63, 67 }
    triangle.Data[7] = []int64{ 99, 65, 04, 28, 06, 16, 70, 92 }
    triangle.Data[8] = []int64{ 41, 41, 26, 56, 83, 40, 80, 70, 33 }
    triangle.Data[9] = []int64{ 41, 48, 72, 33, 47, 32, 37, 16, 94, 29 }
    triangle.Data[10] = []int64{ 53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14 }
    triangle.Data[11] = []int64{ 70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57 }
    triangle.Data[12] = []int64{ 91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48 }
    triangle.Data[13] = []int64{ 63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31 }
    triangle.Data[14] = []int64{ 4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23 }

    for level := 0; level < triangle.Levels; level++ {
        for index := 0; index < len(triangle.data[level]); index++ {
            firstParent, secondParent := triangle.getParentsValues(level, index)

            triangle.Data[level][index] += number.MaxInt64(firstParent, secondParent)
        }
    }

    for _, value := range triangle.Data[levels-1] {
        result = number.MaxInt64(result, value)
    }

    return strconv.FormatInt(result, 10)
}

func init() {
    Registry[18] = problem0018{18, "Find the maximum total from top to bottom of the 15 level triangle."}
}
