package d1

import (
     "testing"
) 

func TestSumCalibration2(t *testing.T) {
    input := []string{
        "two1nine",
        "eightwothree",
        "abcone2threexyz",
        "xtwone3four",
        "4nineeightseven2",
        "zoneight234",
        "7pqrstsixteen", 
    }
    want := 281

    res, err := SumCalibration2(input)
    if want != res || err != nil {
        t.Errorf(`got %v, %v, want match for %v, nil`, res, err, want)
    }

}

type testCase2 struct {
    val string
    want string
}

func TestConvertToNumberString(t *testing.T){

    testCases := []testCase2 {
        {"twothree6vpnvvnshn", "26"},
        {"vvhtbjjrr6xghmzf", "66"},
        {"bfjtdslkdbthree4jvvonezqdthreesrghnnbsix", "36"},
        {"kkbhh5fivepvhzhdsvxvnkrn4", "54"},
        {"4qsqr235twogl21", "41"},
        {"v2seventhreezjfour6", "26"},
        {"38gfqdpkfhdonespqbckbgkkzhgnbqgslkhfl7", "37"},
        {"dtwoneeight6llzcxssgrdfjmjvfbvtwo9", "29"},
        {"sixsixgnpprvjdkgvqmr1", "61"},
        {"7sqthfchpjklpn", "77"},
        {"foureight7hl111rznjfh", "41"},
        {"9six2threefiversbsc", "95"},
        {"8six1ninezjsix", "86"},
        {"6threezjmclknqcztwocfiveninextpdq1", "61"},
        {"gsdsr2seven51", "21"},
        {"one", "11"},
        {"onetwo", "12"},
        {"oneathree", "13"},
    }

    for _, test := range testCases {
        res, err := ConvertToNumberString(test.val)
        if test.want != res || err != nil {
            t.Errorf(`got %v, %v, want match for %v, nil`, res, err, test.want)
        }

    }

}
