package d3

import "fmt"

func SumGearRatio(input []string) (int, error){
    MAX_X := len(input[0])
    MAX_Y := len(input)
    // gear is * adjacent to exactly 2 part numbers
    star := '*'  // 42

    // find all the number coords
    numMap := GetNumberSlots(input)
    fmt.Println(numMap)

    // find all the asterisk coords 
    starMap := GetStarSlots(input)
    fmt.Println(starMap)
    // find all the asterisks, if adjacent to exactly 2 numbers
    // then save as gear. maybe remember the numbers here too?

    stars := make([][]bool, MAX_Y+1)
    for i := range stars {
        stars[i] = make([]bool, MAX_X+1)
    }
    // map all the symbol coords out
    for y, line := range input {
        for x, c := range line {
            if c == star {
                // get adjacent number

                stars[y][x] = true
            }
        }
    }

    adjacents := make([][]bool, MAX_Y+1)
    for i := range adjacents {
        adjacents[i] = make([]bool, MAX_X+1)
    }

    // now go thru the maps and see whats adjacent
    for y, l := range input {
        for x := range l {
            // if the coord is a symbol, flag all "touching" coords
            if stars[y][x] { 
                // self
                adjacents[y][x] = true
                // left and right same line
                adjacents[y][x-1] = true      
                adjacents[y][x+1] = true      

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
    }

    var sum int
    for _, slot := range numMap {
        i := slot.start.x
        e := slot.end.x
        for i < e {
            if adjacents[slot.start.y][i] {
                sum += slot.number 
                break
            }
            i++
        }
    }

    return sum, nil
}

type StarSlot struct {
    pos Coord
}

func GetStarSlots(input []string) ([]StarSlot) {
    var stars []StarSlot

    for y, line := range input {
        for x, c := range line {
            if c == 42 {
                s := StarSlot{Coord{y: y, x:x}}
                stars = append(stars, s)
            }
        }
    }
    return stars
}
