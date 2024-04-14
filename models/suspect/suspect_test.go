package suspect

import (
	"reflect"
	"testing"
)

func TestGetAllSuspects(t *testing.T) {
	suspect1 := Suspect{Name: "Suspect1"}
	suspect2 := Suspect{Name: "Suspect2"}
	suspects = []Suspect{suspect1, suspect2}

	result := getAllSuspects()
	expected := suspects
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("getAllSuspects() = %v; want %v", result, expected)
	}
}

func TestFillSuspects(t *testing.T) {
	suspect1 := Suspect{Name: "Suspect1"}
	suspect2 := Suspect{Name: "Suspect2"}
	suspect3 := Suspect{Name: "Suspect2"}
	filledSuspects := []Suspect{suspect1, suspect2, suspect3}

	fillSuspects(filledSuspects)
	result := suspects
	expected := filledSuspects
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("fillSuspects(filledSuspects) = %v; want %v", result, expected)
	}
}

func TestRemoveAllSuspects(t *testing.T) {
	suspect1 := Suspect{Name: "Suspect3"}
	suspect2 := Suspect{Name: "Suspect4"}
	suspects = []Suspect{suspect1, suspect2}

	removeAllSuspects()
	result := suspects
	expected := []Suspect{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("removeAllSuspects() = %v; want %v", result, expected)
	}

}

func TestAddSuspect(t *testing.T) {
	suspect1 := Suspect{Name: "Suspect1"}

	addSuspect(suspect1)
	result := suspects
	expected := []Suspect{suspect1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("addSuspect(suspect) = %v; want %v", result, expected)
	}
}
