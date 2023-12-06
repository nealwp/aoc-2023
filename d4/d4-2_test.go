package d4

import "testing"

func TestSumScratchCards(t *testing.T) {
    input := []string{
        "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", // 4
        "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", // 2 + 1
        "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", // 2 + 1 + 1
        "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", // 1 + 1 + 1 + 1 
        "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", // 0 + 1     + 1 + 1
        "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", // 0
    }
    want := 30 
    res, err := SumScratchCards(input)
    if res != want || err != nil {
        t.Errorf("got: %v, %v wanted: %v, nil", res, err, want)
    }
}

type testCase struct {
    val string
    want Card
}

func TestGetWinningScratchCards(t *testing.T) {
    tests := []testCase{
        {"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", Card{matches: 4, copies: 1}},
        {"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", Card{matches: 2, copies: 1}},
        {"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", Card{matches: 2, copies: 1}},
        {"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", Card{matches: 1, copies: 1}},
        {"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", Card{matches: 0, copies: 1}}, 
        {"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", Card{matches: 0, copies: 1}},
    }

    for _, test := range tests {
        res := GetWinningScratchCards(test.val)
        if res != test.want {
            t.Errorf("got %v, wanted %v", res, test.want)
        }
    }
}

func TestSetMults(t *testing.T) {
    input := []Card {
        {matches: 4, copies: 1},
        {matches: 2, copies: 1},
        {matches: 2, copies: 1},
        {matches: 1, copies: 1},
        {matches: 0, copies: 1}, 
        {matches: 0, copies: 1},
    }
    
    want := []Card {
        {matches: 4, copies: 1},
        {matches: 2, copies: 2},
        {matches: 2, copies: 4},
        {matches: 1, copies: 8},
        {matches: 0, copies: 14}, 
        {matches: 0, copies: 1},
    }

    res := SetMults(input)
    for i, r := range res {
        if r != want[i] {
            t.Errorf("want %v, got %v", want[i], r)
        }
    }
}
