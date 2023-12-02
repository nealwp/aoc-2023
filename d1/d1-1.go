package d1

import "strconv"

func CalibrationValue(in string) (int, error) {
    if len(in) == 1 {
        i, err := strconv.Atoi(in + in)
        if err != nil {
            panic(err)
        }
        return i, nil
    }

    i, err := strconv.Atoi(in) 
    if err != nil {
        panic(err)
    }
    return i, nil

}



