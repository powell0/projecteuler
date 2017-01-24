package problems

type Problem interface {
    ID() int
    Description() string
    Solve() string
}
