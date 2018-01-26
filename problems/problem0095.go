package problems

import (
    "fmt"
    "github.com/powell0/projecteuler/utilities/math/factoring"
    "strconv"
)

type problem0095 struct {
    id int
    description string
}

func (p problem0095) ID() int {
    return p.id
}

func (p problem0095) Description() string {
    return p.description
}

func (p problem0095) Solve() string {
    const limit = 1000000

    nextInChain := make([]uint64, limit + 1)
    
    for i := uint64(2); i <= limit; i++ {
        factors := factoring.ComputeFactors(i)

        sum := uint64(0)

        for j := 0; j < len(factors)-1; j++ {
            sum += factors[j]
        }

        if sum <= limit {
            nextInChain[i] = sum
        }
    }

    maxChainLength := 0
    maxChainSmallestElement := uint64(0)
    chainMembersSeen := make(map[uint64]struct{})

    for i := 2; i < len(nextInChain); i++ {

        _, ok := chainMembersSeen[uint64(i)]

        if ok {
            continue
        }

        chainLength := 0

        chainMembersSeen[uint64(i)] = struct{}{}

        nextNumber := nextInChain[i]

        for nextNumber != 0 && nextNumber != uint64(i) {
            chainLength++
            chainMembersSeen[nextNumber] = struct{}{}
            nextNumber = nextInChain[nextNumber]
        }

        if nextNumber == uint64(i) {
            if len(chain) > maxChainLength {
                maxChainLength = len(chain)
                maxChainSmallestElement = uint64(i)
            }
        }

        for j := 0; j < len(chain); j++ {
            
        }
    }


    return strconv.FormatUint(maxChainSmallestElement, 10)
}

func init() {
    Registry[95] = problem0095{95, "Find the smallest member of the longest amicable chain with no element exceeding one million."}
}
