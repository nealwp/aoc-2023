package d1

import (
     "testing"
) 

func TestSumCalibration2(t *testing.T) {
    input := []string{
        "two1nine",         // 29
        "eightwothree",     // 83
        "abcone2threexyz",  // 13
        "xtwone3four",      // 24
        "4nineeightseven2", // 42
        "zoneight234",      // 14
        "7pqrstsixteen",    // 76
    }
    want := 281

    res, err := SumCalibration2(input)
    if want != res || err != nil {
        t.Fatalf(`got %v, %v, want match for %v, nil`, res, err, want)
    }

}
