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
    Hands []Handful
}

func SumPossibleGameIds(input string) (int, error) {
    return 0, nil
}

func ParseGameString(input string) (Game, error) {
    var game Game
    s := strings.Split(input, " ")
    // id is always second word in string
    id, err := strconv.Atoi(strings.Split(s[1], ":")[0])
    if err != nil {
        panic(err)
    }
    game.Id = id
    return game, nil 
}
