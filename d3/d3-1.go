package d3

import (
	"fmt"
	"regexp"
	"strconv"
)

func SumSchematic(input []string) (int, error) {
    // char code for period (46)
    // char code string 4 (52) [+48]
    // max for a number = 9 + 48 (57)
    // valid num ranges = (48 - 57)
    // 46 is a period, anything else must be symbol

    // need to know line above and below
    // lines are all fixed width
    // NO SYMBOLS ON TOP LINE OR BOTTOM LINE

    MAX_X := len(input[0])
    MAX_Y := len(input)
    // take each number 
    // scan next char, prev char, next line up, +1, -1, prev line +1, -1
    
    fmt.Println("--symbols--")
    // map all the symbol coords out
    var symbols [][]bool
    for y, line := range input {
        var lineSyms []bool
        for _, c := range line {
            var sc bool 
            if c < 48 && c != 46 {
                sc = true
            }
            lineSyms = append(lineSyms, sc)
        }
        symbols = append(symbols, lineSyms)
        fmt.Println(symbols[y])
    }

    fmt.Println("--numbers--")
    // map all the nums coords out
    var numbers [][]bool
    for y, line := range input {
        var lineNum []bool
        for _, c := range line {
            var num bool 
            if c > 47 && c < 58 {
                num = true
            }
            lineNum = append(lineNum, num)
        }
        numbers = append(numbers, lineNum)
        fmt.Println(numbers[y])
    }

    // setup empty adjacents map
    adjacents := make([][]bool, MAX_Y)
    for i := range adjacents {
        adjacents[i] = make([]bool, MAX_X)
    }

    fmt.Println("--adjacents--")
    // now go thru the maps and see whats adjacent
    for y, l := range input {
        for x := range l {
            // if the coord is a symbol, flag all "touching" coords
            if symbols[y][x] { 
                // left and right same line
                adjacents[y][x-1] = true      
                adjacents[y][x-1] = true      

                // line below 
                adjacents[y+1][x-1] = true      
                adjacents[y+1][x] = true      
                adjacents[y+1][x+1] = true      

                // line above 
                adjacents[y-1][x-1] = true      
                adjacents[y-1][x] = true      
                adjacents[y-1][x+1] = true      
            } 
        }
        fmt.Println(adjacents[y])
    }

    // now collect all the numbers to sum
    type Coord struct {
        y int
        x int
    }

    type NumberSlot struct {
        number int 
        start Coord
        end Coord
    }

    re := regexp.MustCompile("\\d+")
    var numMap []NumberSlot
    for y, line := range input {
        ints := re.FindAllString(line, -1)
        coords := re.FindAllStringIndex(line, -1)
        for i := range ints {
            num, _ := strconv.Atoi(ints[i])
            n := NumberSlot{
                number: num, 
                start: Coord{y: y, x: coords[i][0]},
                end: Coord{y: y, x: coords[i][1]},
            }
            numMap = append(numMap, n)
            fmt.Println(n)
        }
    }

    return 0, nil
}
