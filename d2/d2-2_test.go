package d2

import (
    "testing"
)

func TestSumPowerCubes(t *testing.T) {
    input := []string{
        "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
        "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
        "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
        "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
        "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
    }
    want := 2286
    res, err := SumPowerCubes(input)
    if res != want || err != nil {
        t.Errorf(`got %v, %v, want match for %v, nil`, res, err, want)
    }
}

func TestParseGameString2(t *testing.T) {
    testCases := []testCase{
        {"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 48},
        {"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12},
    }

    for _, i := range testCases {
        res := ParseGameString2(i.test)
        if res != i.want {
            t.Errorf(`got %v, want match for %v`, res, i.want)
        }
    }
}
