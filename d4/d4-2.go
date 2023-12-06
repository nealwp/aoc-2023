package d4

import (
	"fmt"
	"strings"
)

type Card struct {
    matches int
    copies int
}

func SumScratchCards(input []string) (int, error) {
    var sum int
    var cards []Card

    for _, line := range input {
        card := GetWinningScratchCards(line)
        cards = append(cards, card)
    }
    
    mutCards := SetMults(cards)

    fmt.Println(cards)
    for _, card := range mutCards {
        sum += card.copies
    }

    return sum, nil
}

func GetWinningScratchCards(input string) Card {
    var numWinners int
    holderMap := make(map[string]int)

    cards := strings.Split(strings.Split(input, ": ")[1], " | ")
    winners := strings.Split(cards[0], " ")
    holders := strings.Split(cards[1], " ")

    for _, v := range holders {
        if v != "" {
            holderMap[strings.TrimSpace(v)] = 0 
        }
    }

    for _, w := range winners {
        if _, found := holderMap[strings.TrimSpace(w)]; found {
            numWinners += 1
        }
    }

    card := Card{
        matches: numWinners,
        copies: 1,
    }

    return card
}

func SetMults(cards []Card) []Card {
    
    // for number of matches
    // multiply copies * matches for each card

    // bump the number of copies for next cards


    for i, card := range cards {
        m := 1 
        for m <= card.matches {

            if i + m >= len(cards) {
                break
            }

            next := cards[i + m]
            next.copies += 1

            m++
        }


        j := 1
        for j < len(cards) {
            c := cards[j] 
            c.matches = c.copies * c.matches 
            j++
        }
    }
    return cards
}
