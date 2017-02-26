package problems

import (
    "math/big"
)

type problem0071 struct {
    id int
    description string
}

func (p problem0071) ID() int {
    return p.id
}

func (p problem0071) Description() string {
    return p.description
}

func (p problem0071) Solve() string {
    const limit = 1000000
    upperBound := float64(3) / float64(7)

    result := big.NewRat(2, 5)
    lowerBound, _ := result.Float64()

    for d := int64(1); d <= limit; d++ {
        numeratorMin := int64(float64(d) * lowerBound)
        numeratorMax := int64(float64(d) * upperBound)+1

        for n := numeratorMin; n < numeratorMax; n++ {
            tempValue := float64(n) / float64(d)
            
            if  tempValue < upperBound && tempValue > lowerBound {

                result.SetFrac64(n, d)
                lowerBound, _ = result.Float64()
            }
        }
    }

    return result.Num().String()
}

func init() {
    Registry[71] = problem0071{71, "By listing the set of reduced proper fractions for d <= 1,000,000 in ascending order of size, find the numerator of the fraction immediately to the left of 3/7."}
}
