package d1

import "strconv"

func CalibrationValue(in string) (int, error) {
    
    i, err := strconv.Atoi(in + in)
    if err != nil {
        panic(err)
    }

    
    return i, nil
}



