package d1

import (
	"sort"
	"strings"
)

func SumCalibration2(input []string) (int, error) {
    return 0, nil
}

type foundNumStr struct {
    num string
    index int
}

func ConvertToNumberString(input string) (string, error) {

    var found []foundNumStr

    numNames := map[string]string {
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9",
    }

    for k, v := range numNames {
        keyIdx := strings.Index(input, k)
        valIdx := strings.Index(input, v)

        if keyIdx != -1 {
            f := foundNumStr{num: v, index: keyIdx}
            found = append(found, f)
        }

        if valIdx != -1 {
            f := foundNumStr{num: v, index: valIdx}
            found = append(found, f)
        }

    }

    sort.Slice(found, func(i, j int) bool {
        return found[i].index < found[j].index
    }) 

    return found[0].num + found[len(found)-1].num, nil
}



