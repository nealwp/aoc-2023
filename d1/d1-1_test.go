package d1

import (
     "testing"
) 

type testCase struct {
    val string
    want int
}


var testCases = []testCase {
    {"3", 33},
    {"12", 12},
    {"1a3", 13},
    {"a1b", 11},
    {"5lkjsdhga1b", 51},
    {"lkjsd6hgab", 66},
    {"7685", 75},
}


func TestCalibrationValue(t *testing.T) {
    for _, test := range testCases {
        res, err := CalibrationValue(test.val)
        if test.want != res || err != nil {
            t.Fatalf(`got %v, %v, want match for %v, nil`, res, err, test.want)
        }
    }
}

func TestSumCalibration(t *testing.T) {
    input := []string{ "3", "12", "1a3", "a1b"}
    want := 69
    res, err := SumCalibration(input)
    if want != res || err != nil {
        t.Fatalf(`got %v, %v, want match for %v, nil`, res, err, want)
    }
}
