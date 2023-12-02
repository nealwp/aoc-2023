package d2

import (
	"strconv"
	"strings"
)

type Handful struct {
    Blue int
    Red int
    Green int
}

type Game struct {
    Id int
    Possible bool
}

func SumPossibleGameIds(input []string) (int, error) {
    var sum int
    for _, s := range input {
        sum += ParseGameString(s)
    }

    return sum, nil
}

func ParseGameString(input string) (int) {
    var game Game
    game.Possible = true

    // id is always second word in string
    s := strings.Split(input, ":")
    gameStr := s[0]

    id, err := strconv.Atoi(strings.Split(gameStr, " ")[1])
    if err != nil {
        panic(err)
    }
    game.Id = id

    handsStr := s[1]
    hands := strings.Split(handsStr, ";")
    
    for _, h := range hands {
        var hand Handful
        for _, e := range strings.Split(h, ",") {
            blueIdx := strings.Index(e, "blue")
            if blueIdx != -1 {
                blueNum := strings.TrimSpace(strings.Split(e, "blue")[0])
                hand.Blue, err = strconv.Atoi(blueNum)
            }

            redIdx := strings.Index(e, "red")
            if redIdx != -1 {
                redNum := strings.TrimSpace(strings.Split(e, "red")[0])
                hand.Red, err = strconv.Atoi(redNum)
            }

            greenIdx := strings.Index(e, "green")
            if greenIdx != -1 {
                greenNum := strings.TrimSpace(strings.Split(e, "green")[0])
                hand.Green, err = strconv.Atoi(greenNum)
            }
        }

        if hand.Blue > 14 || hand.Red > 12 || hand.Green > 13 {
            game.Possible = false
            break
        }
    }

    if game.Possible {
        return game.Id
    } else {
        return 0
    }
}
