package vttimetable

//TODO: should be moved to its own package

type errorString struct {
	str string
}

/*GenericError generates a new generic error with the string passed in*/
func GenericError(text string) error {
	return &errorString{text}
}

/*ParsingError is an error string specific to parsing.*/
func ParsingError(text string) error {
	return &errorString{"Parsing Error: " + text}
}

func (e *errorString) Error() string {
	return e.str
}
