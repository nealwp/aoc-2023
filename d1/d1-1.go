package d1

import "strconv"

func CalibrationValue(input string) (int, error) {

    var numStr []string

    for _, char := range input {
        if _, err := strconv.Atoi(string(char)); err == nil {
            numStr = append(numStr, string(char))
        }
    }

    if len(numStr) == 1 {
        i, err := strconv.Atoi(numStr[0] + numStr[0]) 
        if err != nil {
            panic(err)
        }
        return i, nil
    } else {
        i, err := strconv.Atoi(numStr[0] + numStr[1])
        if err != nil {
            panic(err)
        }
        return i, nil
    }

}



