package swearjar

import (
	"reflect"
	"testing"
)

func TestLoadSwears(t *testing.T) {
	swears, err := Load()

	if err != nil {
		t.Error(err)
	}

	if swears == nil {
		t.Error("Swears could not be loaded")
	}
}

func TestLoadSwearsWithNonexistentFile(t *testing.T) {
	swears, err := Load("nonexistent.json")

	if err == nil {
		t.Error(err)
	}

	if swears != nil {
		t.Error("Swears could be loaded when they should not")
	}
}

func TestLoadSwearsWithBadJSON(t *testing.T) {
	swears, err := Load("swearjar_test_bad.json")

	if err == nil {
		t.Error(err)
	}

	if swears != nil {
		t.Error("Swears could be loaded when they should not")
	}
}

func TestProfane(t *testing.T) {
	swears, err := Load()

	if err != nil {
		t.Error(err)
	}

	profane, err := swears.Profane("fuck is a word")

	if err != nil {
		t.Error(err)
	}

	if profane == false {
		t.Error("fuck is profane")
	}
}

func TestScorecard(t *testing.T) {
	swears, err := Load()

	if err != nil {
		t.Error(err)
	}

	profane, reasons, err := swears.Scorecard("what an asslick")

	if err != nil {
		t.Error(err)
	}

	if profane == false {
		t.Error("asslick is profane")
	}

	expectedReasons := []string{"insult"}
	if !reflect.DeepEqual(reasons, expectedReasons) {
		t.Error("does not match reason")
	}
}

func TestScorecardNoMatch(t *testing.T) {
	swears, err := Load()

	if err != nil {
		t.Error(err)
	}

	profane, reasons, err := swears.Scorecard("this is a lovfuckely sentshitence with no foul languassge")

	if err != nil {
		t.Error(err)
	}

	if profane == true {
		t.Error("this is not profane")
	}

	if len(reasons) > 0 {
		t.Error("should not have a reason")
	}
}
