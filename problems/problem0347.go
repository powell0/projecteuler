package problems

import (
    "github.com/powell0/projecteuler/utilities/math/primes"
    "strconv"
)

type problem0347 struct {
    id int
    description string
}

func (p problem0347) ID() int {
    return p.id
}

func (p problem0347) Description() string {
    return p.description
}

func (p problem0347) Solve() string {
    const limit = 10000000
    count := 0

    primesList := primes.ComputePrimes(limit/2)

    count = len(primesList)

    return strconv.Itoa(count)
}

func init() {
    Registry[347] = problem0347{347, "For two distinct primes p and q let M(p,q,N) be the largest positive integer <= N only divisible by both p and q and M(p,q,N)=0 if such a positive integer does not exist.  Let S(N) be the sum of all distinct M(p,q,N). Find S(10000000)."}
}
