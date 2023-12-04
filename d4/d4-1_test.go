package d4

import "testing"

func TestSomeFunction(t *testing.T) {
    input := []string{}
    want := 0
    res, err := SomeFunction(input)
    if res != want || err != nil {
        t.Fatalf("got: %v, %v wanted: %v, nil", res, err, want)
    }
}
