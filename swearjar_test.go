package swearjar

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadSwears(t *testing.T) {
	swears, err := loadSwears()

	if err != nil {
		t.Error(err)
	}

	if swears == nil {
		t.Error("Swears is nil")
	}

}

func TestProfane(t *testing.T) {
	profane, err := Profane("fuck is a word")

	if profane == false {
		t.Error("fuck is profane")
	}

	if err != nil {
		t.Error(err)
	}
}

func TestScorecard(t *testing.T) {
	profane, reasons, err := Scorecard("what an asslick")

	if profane == false {
		t.Error("asslick is profane")
	}

	expectedReasons := []string{"insult"}
	fmt.Printf("%+v - %+v\n", reasons, expectedReasons)
	if !reflect.DeepEqual(reasons, expectedReasons) {
		t.Error("does not match reason")
	}

	if err != nil {
		t.Error(err)
	}
}
