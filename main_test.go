package main

import (
	"os"
	"testing"
)

/*TestReadWithAdditionalTimes tests reading with a file that contains a row that contains "Additional Times"*/
func TestReadWithAdditionalTimes(t *testing.T) {
	reader, err := os.Open("testinputs/foo1.txt")
	if err != nil {
		t.Errorf("Ran into an error while attempting to read file")
	}
	_, err = ReadInput(reader)
	if err != nil {
		t.Errorf("Ran into an error while attempting to read file")
	}
}
