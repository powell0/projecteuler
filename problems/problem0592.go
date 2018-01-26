package problems

import (
    "fmt"
    "math/big"
    //"strconv"
)

type problem0592 struct {
    id int
    description string
}

func (p problem0592) ID() int {
    return p.id
}

func (p problem0592) Description() string {
    return p.description
}

func (p problem0592) Solve() string {
    const limit = 11
    /*
    const tenThousand = 10000

    number := uint64(1)

    for i := uint64(2); i <= limit; i++ {
        number *= i
    }

    result := uint64(1)

    for i := uint64(2); i <= number; i++ {
        product := i //% tenThousand

        if product == 0 {
            continue
        }

        result *= product

        remainder := result % 10

        for remainder == 0 {
            result /= 10

            remainder = result % 10
        }

        result %= tenThousand
    }

    return strconv.FormatUint(result, 10)*/

    number := int64(1)

    for i := int64(2); i <= limit; i++ {
        number *= i
    }

    
    zero := big.NewInt(0)
    ten := big.NewInt(10)
    //tenThousand := big.NewInt(10000)
    digitLimit := new(big.Int).Exp(ten, big.NewInt(19), nil)
    //fmt.Println(digitLimit.String())
    result := big.NewInt(1)

    for i := int64(2); i <= number; i++ {
        product := big.NewInt(i)
        //product.Mod(product, tenThousand)

        result.Mul(result, product)

        remainder := new(big.Int).Mod(result, ten)

        for remainder.Cmp(zero) == 0 {
            result.Div(result, ten)
            remainder = new(big.Int).Mod(result, ten)
        }

        result.Mod(result, digitLimit)
        
    }

    fmt.Println(result.String())

    result.Mod(result, big.NewInt(10000))
    
    return result.String()
    
}


func init() {
    Registry[592] = problem0592{592, "For any N, let f(N) be the last twelve hexadecimal digits before the trailing zeroes in N!.  Find f(20!)."}
}
