package d3

func SumGearRatio(input []string) (int, error){
    MAX_X := len(input[0])
    MAX_Y := len(input)
    // gear is * adjacent to exactly 2 part numbers
    star := '*' // 42

    // find all the number coords
    numMap := GetNumberSlots(input)

    var sum int

    // find all the asterisks, if adjacent to exactly 2 numbers
    // then save as gear. maybe remember the numbers here too?
    stars := make([][]bool, MAX_Y+1)
    for i := range stars {
        stars[i] = make([]bool, MAX_X+1)
    }

    // map all the star coords out
    for y, line := range input {
        for x, c := range line {
            if c == star {
                // get all nums adjacent to this star
                starCoord := Coord{y: y, x: x}
                adjNums := GetAdjacentNumbers(starCoord, numMap)
                if len(adjNums) == 2 {
                    sum += (adjNums[0] * adjNums[1])
                }
            }
        }
    }

    return sum, nil
}

func GetAdjacentNumbers(star Coord, numMap []NumberSlot) ([]int) {
    var adjNums []int
    for _, slot := range numMap {
        // only scan lines 1 up, 1 down, or same
        if slot.start.y > (star.y + 1) || slot.start.y < (star.y - 1) {
            continue
        }
        i := slot.start.x
        e := slot.end.x
        for i < e {
            // left or right same line
            if i == star.x + 1 && slot.start.y == star.y || i == star.x - 1 && slot.start.y == star.y {
                adjNums = append(adjNums, slot.number)
                break
            }

            // left, centered, or right up line 
            if i == star.x + 1 && slot.start.y == star.y - 1 || i == star.x && slot.start.y == star.y - 1 || i == star.x - 1 && slot.start.y == star.y -1 {
                adjNums = append(adjNums, slot.number)
                break
            }


            // left, centered, or right down line 
            if i == star.x + 1 && slot.start.y == star.y + 1 || i == star.x && slot.start.y == star.y + 1 || i == star.x - 1 && slot.start.y == star.y + 1 {
                adjNums = append(adjNums, slot.number)
                break
            }

            i++
        }
    }
    return adjNums
}
