package prose

import (
	"testing"
)

func TestJoinWithCommas(t *testing.T) {

	input := []string{"apple", "orange", "grape"}
	want := "apple, orange and grape"

	get := JoinWithCommas(input)
	// fmt.Println(get)

	if get != want {
		t.Errorf("Input :: %#v", input)
		t.Errorf("Wanted :: %#v", want)
		t.Errorf("Output :: %#v", get)
	}

}

// 수정 필요

// type testData struct {
// 	Input []string
// 	Want  string
// }

// func TestJoinWithCommasArray(t *testing.T) {

// 	testList := []testData{
// 		testData{Input: []string{}, Want: ""},
// 		testData{Input: []string{"apple", "orange", "grape"}, Want: "apple, orange and grape"},
// 	}

// 	// input := []string{"apple", "orange", "grape"}
// 	// want := "apple, orange and grape"

// 	for _, data := range testList {
// 		get := JoinWithCommas(data.Input)

// 		if get != data.Want {
// 			t.Errorf("Input :: %#v", data.Input)
// 			t.Errorf("Wanted :: %#v Output :: %#v", data.Want, get)
// 		}
// 	}

// }
