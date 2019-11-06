package vttimetable

import (
	"os"
	"testing"
)

/*TestReadWithAdditionalTimes tests reading with a file that contains a row that contains "Additional Times"*/
func TestReadWithAdditionalTimes(t *testing.T) {
	reader, err := os.Open("testinputs/foo1.txt")
	if err != nil {
		t.Errorf("Ran into an error while attempting to open file")
	}
	_, err = ReadInput(reader)
	if err != nil {
		t.Errorf("Ran into an error while attempting to read file contents")
	}
}

func TestReadRequestToDB(t *testing.T) {
	cq := CreateEmptyQuery().SetCampus(CampusCodes.BLACKSBURG).SetCrn("828").SetTermYear(Term("2019", FALL))
	resp, err := SendQuery(timetableURL, cq)
	if err != nil {
		t.Errorf("Ran into an error while attempting to send request")
	}
	_, err = ReadInput(resp.Body)
	if err != nil {
		t.Errorf("Ran into an error while attempting to read response")
	}
}
