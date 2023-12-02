package d1

import (
     "testing"
) 

func TestCalibrationValue(t *testing.T) {
    value := "3"
    want := 33 
    res, err := CalibrationValue(value)
    if want != res || err != nil {
        t.Fatalf(`got %v, %v, want match for %v, nil`, res, err, want)
    }
}

