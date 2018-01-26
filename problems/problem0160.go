package problems

import (
    //"github.com/powell0/projecteuler/utilities/math/factorial"
    //"github.com/powell0/projecteuler/utilities/math/number"
    //"github.com/powell0/projecteuler/utilities/math/primes"
    "fmt"
    "math/big"
    "strconv"
)

type problem0160 struct {
    id int
    description string
}

func (p problem0160) ID() int {
    return p.id
}

func (p problem0160) Description() string {
    return p.description
}

func (p problem0160) Solve() string {

    const target = 1000000000
    const digitLimit = 100000
    //const target = 1000000000000
    //const digitLimit = 100000


    result := ComputeRecursively2(target, 0, digitLimit)

    zero := big.NewInt(0)
    ten := big.NewInt(10)
    bob := big.NewInt(7)
    //bob := new(big.Int).MulRange(1,target)

    remainder := new(big.Int).Mod(bob, ten)

    for remainder.Cmp(zero) == 0 {
        bob.Div(bob, ten)
        remainder.Mod(bob, ten)
    }

    text := bob.String()

    fmt.Println(result)
    fmt.Println(text[len(text)-1:])

    return strconv.FormatUint(target, 10)
}

func ComputeRecursively2(a uint64, b uint64, mod uint64) uint64 {
    d := a - b
    result := uint64(0)

    if d < 0 {
        result = 0
    } else if d == 0 {
        result = 1
    } else if d == 1 {
        result = a

        for result % 10 == 0 {
            result /= 10
        }
    } else if d == 2 {
        result = a * (a - 1)

        for result % 10 == 0 {
            result /= 10
        }
    } else if d == 3 {
        result = a * (a - 1) * (a - 2)

        for result % 10 == 0 {
            result /= 10
        }
    } else {
        m := (a + b) / 2

        factor1 := ComputeRecursively2(a, m, mod)
        factor2 := ComputeRecursively2(m, b, mod)

        result = factor1*factor2

        for result % 10 == 0 {
            result /= 10
        }
    }

    result %= mod
    return result
}

func ComputeRecursively(a uint64, b uint64) uint64 {
    d := a - b
    m := (a + b) / 2

    if d < 0 {
        return 0
    } else if d == 0 {
        return 1
    } else if d == 1 {
        return a
    } else if d == 2 {
        return a * (a - 1)
    } else if d == 3 {
        return a * (a - 1) * (a - 2)
    } else {
        return ComputeRecursively(a, m)*ComputeRecursively(m, b)
    }
}

func init() {
    Registry[160] = problem0160{160, "For any N, let f(N) be the last five digits before the trailing zeroes in N!.  Find f(1,000,000,000,000)."}
}
