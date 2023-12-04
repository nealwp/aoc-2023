package d4

import (
	"fmt"
	"strings"
)

func SumScratchCardWinnings(input []string) (int, error) {
    var sum int
    for _, line := range input {
        sum += GetWinnings(line)
    }
    return sum, nil
}

func GetWinnings(input string) int {
    var points int
    holderMap := make(map[string]int)

    cards := strings.Split(strings.Split(input, ": ")[1], " | ")
    winners := strings.Split(cards[0], " ")
    holders := strings.Split(cards[1], " ")

    for _, v := range holders {
        if v != "" {
            holderMap[strings.TrimSpace(v)] = 0 // this really doesn't matter
        }
    }

    fmt.Println(holderMap)
    for _, w := range winners {
        if _, found := holderMap[strings.TrimSpace(w)]; found {
            if points == 0 {
                points += 1
            } else {
                points *= 2
            }
        }
    }

    return points
}


