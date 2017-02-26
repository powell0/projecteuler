package problems

import (
    "strconv"
)

type problem0073 struct {
    id int
    description string
}

func (p problem0073) ID() int {
    return p.id
}

func (p problem0073) Description() string {
    return p.description
}

func (p problem0073) Solve() string {
    const limit = 12000
    upperBound := float64(1) / float64(2)
    lowerBound := float64(1) / float64(3)

    fractions := make(map[float64]struct{})

    for d := int64(1); d <= limit; d++ {
        numeratorMin := int64(float64(d) * lowerBound)
        numeratorMax := int64(float64(d) * upperBound)+1

        for n := numeratorMin; n < numeratorMax; n++ {
            tempValue := float64(n) / float64(d)
            
            if  tempValue < upperBound && tempValue > lowerBound {
                fractions[tempValue] = struct{}{}
            }
        }
    }

    return strconv.Itoa(len(fractions))
}

func init() {
    Registry[73] = problem0073{73, "How many fractions lie between 1/3 and 1/2 in the sorted set of reduced proper fractions for d <= 12,000?"}
}
