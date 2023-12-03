package d3

import "fmt"

func SumSchematic(input []string) (int, error) {
    // char code for period (46)
    // char code string 4 (52) [+48]
    // max for a number = 9 + 48 (57)
    // valid num ranges = (48 - 57)
    // 46 is a period, anything else must be symbol

    // need to know line above and below
    // lines are all fixed width

    //MAX_LEN := len(input[0])

    // take each number
    // scan next char, prev char, next line up, +1, -1, prev line +1, -1
    
    // map all the symbol coords out
    var symLines [][]bool
    for _, line := range input {
        var lineSyms []bool
        for _, c := range line {
            var sc bool 
            if c < 48 && c != 46 {
                sc = true
            }
            lineSyms = append(lineSyms, sc)
        }
        symLines = append(symLines, lineSyms)
    }

    // map all the nums coords out
    var numLines [][]bool
    for _, line := range input {
        var lineNum []bool
        for _, c := range line {
            var num bool 
            if c > 47 && c < 58 {
                num = true
            }
            lineNum = append(lineNum, num)
        }
        numLines = append(numLines, lineNum)
    }

    // now go thru the maps and see whats adjacent
    for i, l := range input {
        for ii := range l {
            if numLines[i][ii] {
                if symLines[i][ii] || symLines[i][ii - 1] || symLines[i][ii + 1] {
                }
            }
        }
    }
    fmt.Println(symLines)
    fmt.Println(numLines)
    
    return 0, nil
}
