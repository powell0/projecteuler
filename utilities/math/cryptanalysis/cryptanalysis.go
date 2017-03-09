package cryptanalysis

import (
    "sort"
)

var asciiCharacterFrequency []byte

type symbolCountPair struct {
    symbol byte
    count int
}

func AnalyzeSubstitutionCipher(cipherTextSymbolCounts map[byte]int) map[byte]byte {
    cipherTextSymbols := make([]symbolCountPair, 0, len(cipherTextSymbolCounts))

    for symbol, count := range cipherTextSymbolCounts {
        cipherTextSymbols = append(cipherTextSymbols, symbolCountPair{symbol, count})
    }

    sort.Slice(cipherTextSymbols, func(i int, j int) bool { return cipherTextSymbols[i].count > cipherTextSymbols[j].count })

    substitutions := make(map[byte]byte)

    for i := 0; i < len(cipherTextSymbols) && i < len(asciiCharacterFrequency); i++ {
        substitutions[cipherTextSymbols[i].symbol] = asciiCharacterFrequency[i]
    }

    return substitutions
}

func init() {
    asciiCharacterFrequency = []byte{
        32,     // space
        101,    // e
        116,    // t
        97,     // a
        111,    // o
        105,    // i
        110,    // n
        115,    // s
        114,    // r
        104,    // h
        108,    // l
        100,    // d
        99,     // c
        117,    // u
        109,    // m
        102,    // f
        103,    // g
        112,    // p
        121,    // y
        119,    // w
        98,     // b
        44,     // ,
        46,     // .
        118,    // v
        107,    // k
        45,     // -
        34,     // "
        95,     // _
        39,     // '
        120,    // x
        41,     // )
        40,     // (
        59,     // ;
        48,     // 0
        106,    // j
        49,     // 1
        113,    // q
        61,     // =
        50,     // 2
        58,     // :
        122,    // z
        47,     // /
        42,     // *
        33,     // !
        63,     // ?
        36,     // $
        51,     // 3
        53,     // 5
        62,     // >
        123,    // {
        125,    // }
        52,     // 4
        57,     // 9
        91,     // [
        93,     // ]
        56,     // 8
        54,     // 6
        55,     // 7
        92,     // /
        43,     // +
        124,    // |
        38,     // &
        60,     // <
        37,     // %
        64,     // @
        35,     // #
        94,     // ^
        96,     // `
        126,    // ~
        0}
}
