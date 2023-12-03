package d3

import (
    "testing"
)

func TestSomething(t *testing.T) {
    input := []string{
        "467..114..",
        "...*......",
        "..35..633.",
        "......#...",
        "617*......",
        ".....+.58.",
        "..592.....",
        "......755.",
        "...$.*....",
        ".664.598..",
    }


    want := 4361
    res, err := SumSchematic(input)
    if res != want || err != nil {
        t.Errorf(`got %v, %v, want match for %v, nil`, res, err, want)
    }
}
