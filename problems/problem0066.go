package problems

import (
    "github.com/powell0/projecteuler/utilities/math/continuedfractions"
    "math/big"
    "strconv"
)

type problem0066 struct {
    id int
    description string
}

func (p problem0066) ID() int {
    return p.id
}

func (p problem0066) Description() string {
    return p.description
}

func (p problem0066) Solve() string {
    const limit = 1000

    maxX := big.NewInt(0)
    bestD := int64(0)

    for D := int64(1); D <= limit; D++ {
        representation := continuedfractions.ComputeSquareRootRepresenation(D)

        if len(representation) > 1 {
            x, _ := findPellEquationFundamentalSolution(D, representation)


            if x.Cmp(maxX) > 0 {
                maxX = x
                bestD = D
            }
        }
    }

    return strconv.FormatInt(bestD, 10)
}

func findPellEquationFundamentalSolution(D int64, continuedFractionRepresentation []int64) (x *big.Int, y *big.Int) {
    expansions := 0
    rep := continuedFractionRepresentation[1:]
    repLength := len(rep)
    a0 := big.NewRat(continuedFractionRepresentation[0], 1)
    coefficients := make([]int64, 0, 10)

    for true {
        index := expansions % repLength
        coefficients = append([]int64{rep[index]}, coefficients...)

        result := big.NewRat(1, coefficients[0])

        for i := 1; i < len(coefficients); i++ {
            coefficient := big.NewRat(coefficients[i], 1)
            result.Add(coefficient, result)
            result.Inv(result)
        }

        result.Add(a0, result)

        x = result.Num()
        y = result.Denom()

        if isPellEquationSolution(big.NewInt(D), x, y) {
            break
        }

        expansions++
    }

    return
}

func isPellEquationSolution (D *big.Int, x *big.Int, y *big.Int) bool {
    xSquare := new(big.Int).Mul(x, x)
    ySquare := new(big.Int).Mul(y, y)
    result := new(big.Int).Sub(xSquare, new(big.Int).Mul(D, ySquare))
    
    one := big.NewInt(1)

    return result.Cmp(one) == 0
}

func init() {
    Registry[66] = problem0066{66, "Find the value of D <= 1000 in minimal solutions of x where x^2 - Dy^2 = 1 for which the largest value of x is obtained."}
}
