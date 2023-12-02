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
}


func TestCalibrationValue(t *testing.T) {
    for _, test := range testCases {
        res, err := CalibrationValue(test.val)
        if test.want != res || err != nil {
            t.Fatalf(`got %v, %v, want match for %v, nil`, res, err, test.want)
        }
    }
}

