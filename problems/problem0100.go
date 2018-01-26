package problems

import (
    "fmt"
    //"github.com/powell0/projecteuler/utilities/math/factoring"
    //"github.com/powell0/projecteuler/utilities/math/factorial"
    //"github.com/powell0/projecteuler/utilities/math/number"
    //"github.com/powell0/projecteuler/utilities/math/primes"
    "math/big"
)

type problem0100 struct {
    id int
    description string
}

func (p problem0100) ID() int {
    return p.id
}

func (p problem0100) Description() string {
    return p.description
}

func (p problem0100) Solve() string {
    const start = 1000000000000
    one := big.NewInt(1)
    two := big.NewInt(2)
    result := big.NewInt(0)
    i := int64(start)

    for true {
        if i % 1000000 == 0 {
            fmt.Println(i)
        }

        if i < start {
            fmt.Println("OverFlow")
            break
        }

        i++
        T := new(big.Int).Mul(big.NewInt(i), big.NewInt(i-1))
        T.Mul(T, two)
        T.Add(T, one)

        R := new(big.Int).Sqrt(T)

        RR := new(big.Int).Mul(R, R)

        D := new(big.Int).Mod(R, two)
        
        // If square root of T is an integer and odd
        if T.Cmp(RR) == 0 && D.Cmp(one) == 0 {
            R.Add(R, one)
            R.Div(R, two)
            result.Set(R)
            fmt.Println(i)
            break
        }

    }

    return result.String()
}

func init() {
    Registry[100] = problem0100{100, "By finding the first arrangement of having a 50% chance of picking two blue discs in a row to contain over 1,000,000,000,000 discs in total, determine the number of blue discs that the box would contain."}
}
