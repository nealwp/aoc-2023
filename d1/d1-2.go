package d1

import (
	"strconv"
)

func SumCalibration2(input []string) (int, error) {
    sum := 0
    for _, i := range input {
        o, err := ConvertToNumberString(i)    
        if err != nil {
            panic(err)
        }
        num, err := strconv.Atoi(o)
        if err != nil {
            panic(err)
        }
        sum += num
    }

    return sum, nil
}

func ConvertToNumberString(input string) (string, error) {

    numWords := map[string]string {
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9",
        "1": "1",
        "2": "2",
        "3": "3",
        "4": "4",
        "5": "5",
        "6": "6",
        "7": "7",
        "8": "8",
        "9": "9",
    }

    var res string

    // for each letter in input 
    for i := range input {

        // go thru the map
        for k, v := range numWords {

            // skip if word is longer than remaining string
            if len(k) > len(input[i:]){
                continue
            }
            
            // window is length of word we looking for
            // get the characters under the window
            strWind := string(input[i:i + len(k)])

            // if match, append value to result and get outta there
            if strWind == k {
                res += v
                break
            }
        }
    }

    return string(res[0]) + string(res[len(res)-1]), nil
}

