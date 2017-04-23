package problems

import (
    "github.com/powell0/projecteuler/utilities/bits"
    "github.com/powell0/projecteuler/utilities/math/number"
    "math"
    "math/big"
    "strconv"
)

type problem0121 struct {
    id int
    description string
}

func (p problem0121) ID() int {
    return p.id
}

func (p problem0121) Description() string {
    return p.description
}

func (p problem0121) Solve() string {
    const rounds = 15
    start := 0
    end := number.PowInt(2, rounds) - 1
    drawOdds := make([][]*big.Rat, 2)
    drawOdds[0] = make([]*big.Rat, rounds) // chance of red disc
    drawOdds[1] = make([]*big.Rat, rounds) // chance of blue disc

    redDiscs := int64(1)
    for i := 0; i < rounds; i++ {
        drawOdds[0][i] = big.NewRat(redDiscs, redDiscs + 1)
        drawOdds[1][i] = big.NewRat(1, redDiscs + 1)
        redDiscs++
    }

    houseWinChance := big.NewRat(1, 1)
    houseLoseChance := big.NewRat(0, 1)

    for n := start; n <= end; n++ {
        count := bits.CountSetBits(uint64(n))

        if count > rounds / 2 {
            lossProbability := HouseLossProbability(uint(n), drawOdds, rounds)

            houseLoseChance.Add(houseLoseChance, lossProbability)
        }
    }

    houseWinChance.Sub(houseWinChance, houseLoseChance)

    breakEvenPot := new(big.Rat).Quo(houseWinChance, houseLoseChance)

    allocatedMoney, _ := breakEvenPot.Float64()
    allocatedMoney = math.Ceil(allocatedMoney)
    
    return strconv.Itoa(int(allocatedMoney))
}

func HouseLossProbability(drawPattern uint, drawOdds [][]*big.Rat, rounds int) *big.Rat {
    probability := big.NewRat(1, 1)

    for r := 0; r < rounds; r++ {
        disc := drawPattern & 0x1
        
        probability.Mul(probability, drawOdds[disc][r])

        drawPattern >>= 1
    }

    return probability
}

func init() {
    Registry[121] = problem0121{121, "Find the maximum prize fund that should be allocated to a single game in which fifteen turns are played."}
}
