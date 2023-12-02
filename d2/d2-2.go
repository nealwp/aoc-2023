package d2

import (
    "strings"
    "strconv"
)

func SumPowerCubes(input []string) (int, error) {
    var sum int
    for _, i := range input {
        sum += ParseGameString2(i)
    }
    return sum, nil
}

func ParseGameString2(input string) int {

    s := strings.Split(input, ":")
    handsStr := s[1]
    hands := strings.Split(handsStr, ";")
    
    minRed := 0
    minGreen := 0
    minBlue := 0

    for _, h := range hands {
        for _, e := range strings.Split(h, ",") {
            blueIdx := strings.Index(e, "blue")
            if blueIdx != -1 {
                blueStr := strings.TrimSpace(strings.Split(e, "blue")[0])
                blueNum, _ := strconv.Atoi(blueStr)
                if blueNum > minBlue {
                    minBlue = blueNum
                }
            }

            redIdx := strings.Index(e, "red")
            if redIdx != -1 {
                redStr := strings.TrimSpace(strings.Split(e, "red")[0])
                redNum, _ := strconv.Atoi(redStr)
                if redNum > minRed {
                    minRed = redNum
                }
            }

            greenIdx := strings.Index(e, "green")
            if greenIdx != -1 {
                greenStr := strings.TrimSpace(strings.Split(e, "green")[0])
                greenNum, _ := strconv.Atoi(greenStr)
                if greenNum > minGreen {
                    minGreen = greenNum
                }
            }
        }
    }

    return minBlue * minGreen * minRed
}
