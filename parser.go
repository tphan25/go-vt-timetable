package main

import (
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

/*Course represents a single course as read from the HTTP response that we are parsing.*/
type Course struct {
	Crn             string
	CourseName      string
	ClassName       string
	SectionType     string
	CreditHours     string
	Capacity        string
	Instructor      string
	Days            string
	TimeBegin       string
	TimeEnd         string
	Location        string
	ExamTime        string
	AdditionalTimes AdditionalTimes
}

/*AdditionalTimes is an additional struct, optional for classes that specifically state additional times.*/
type AdditionalTimes struct {
	Days      string
	TimeBegin string
	TimeEnd   string
	Location  string
}

/*ReadInput takes in an input of an HTML file corresponding to the VT timetable page. It parses and returns a list of courses.*/
func ReadInput(reader io.Reader) ([]Course, error) {
	tokenizer := html.NewTokenizer(reader)

	courseList := make([]Course, 0)

	for {
		next := tokenizer.Next()

		switch {
		//html.ErrorToken can indicate EOF, or we return an error as well
		case next == html.ErrorToken:
			return courseList, tokenizer.Err()
		case next == html.StartTagToken:
			token := tokenizer.Token()

			for _, a := range token.Attr {
				if a.Key == "class" && a.Val == "dataentrytable" {
					//Skips through header row of table because last element in header row contains "Exam"
					err := skipToText(tokenizer, "Exam")
					if err != nil {
						log.Panic(err)
					}
					//Read through table until no more <tr> elements
					for currCourse, hasNext := readCourseRow(tokenizer); hasNext; currCourse, hasNext = readCourseRow(tokenizer) {
						courseList = append(courseList, currCourse)
					}
					return courseList, err
				}
			}
		}

	}
}

/*skipToText reads until an html.TextToken equates to str or EOF.*/
func skipToText(tokenizer *html.Tokenizer, str string) error {
	for {
		next := tokenizer.Next()

		switch {
		case next == html.ErrorToken:
			return tokenizer.Err()
		case next == html.TextToken:
			token := tokenizer.Token()

			if token.Data == str {
				return nil
			}
		}
	}
}

//Skips ending tag tokens as well
func getNextStartingTag(tokenizer *html.Tokenizer) (html.Token, error) {
	for {
		next := tokenizer.Next()
		switch {
		case next == html.ErrorToken:
			return html.Token{}, tokenizer.Err()
		case next == html.StartTagToken:
			return tokenizer.Token(), nil
		}
	}
}

//Gets text from tokenizer when on starting tag (i.e. <b>Hello</b> returns Hello)
func getNextText(tokenizer *html.Tokenizer) (html.Token, error) {
	for {
		next := tokenizer.Next()
		switch {
		case next == html.ErrorToken:
			return html.Token{}, tokenizer.Err()
		case next == html.TextToken:
			return tokenizer.Token(), nil
		}
	}
}

func getTextFieldsFromTokenizer(tokenizer *html.Tokenizer, tagsBefore []string, fieldName string) (string, error) {
	/*
		<Starting Token>
		<TD>
			<P>
			<A>
				<B>field to find</B>
		Assuming the tokenizer starts at <Starting Token>,  if we pass in ["TD", "P", "A", "B"],
		our tokenizer will parse through all of the tokens until it reaches B then reach inside for the text
		between tags for B (or whatever TextToken follows B)
	*/
	errorMsg := "In parsing for " + fieldName + " - "

	for _, str := range tagsBefore {
		nextTag, err := getNextStartingTag(tokenizer)
		if err != nil {
			return "", err
		}
		if nextTag.Data != str {
			return "", ParsingError(errorMsg + "Token returned should be " + str + ", but got " + nextTag.Data)
		}
	}

	text, err := getNextText(tokenizer)
	return strings.TrimSpace(text.Data), err
}

func readCourseRow(tokenizer *html.Tokenizer) (Course, bool) {
	currCourse := Course{}

	nextTag, err := getNextStartingTag(tokenizer)
	if err != nil {
		log.Panic(err)
	}
	//No more <tr> tags
	if nextTag.Data != "tr" {
		return currCourse, false
	}

	crn, err := getTextFieldsFromTokenizer(tokenizer, []string{"td", "p", "a", "b"}, "CRN")
	if err != nil {
		log.Panic(err)
	}
	currCourse.Crn = crn

	courseName, err := getTextFieldsFromTokenizer(tokenizer, []string{"td", "font"}, "courseName")
	if err != nil {
		log.Panic(err)
	}
	currCourse.CourseName = courseName

	className, err := getTextFieldsFromTokenizer(tokenizer, []string{"td"}, "className")
	if err != nil {
		log.Panic(err)
	}
	currCourse.ClassName = className

	sectionType, err := getTextFieldsFromTokenizer(tokenizer, []string{"td", "p"}, "sectionType")
	if err != nil {
		log.Panic(err)
	}
	currCourse.SectionType = sectionType

	creditHours, err := getTextFieldsFromTokenizer(tokenizer, []string{"td", "p"}, "creditHours")
	if err != nil {
		log.Panic(err)
	}
	currCourse.CreditHours = creditHours

	capacity, err := getTextFieldsFromTokenizer(tokenizer, []string{"td"}, "capacity")
	if err != nil {
		log.Panic(err)
	}
	currCourse.Capacity = capacity

	instructor, err := getTextFieldsFromTokenizer(tokenizer, []string{"td"}, "instructor")
	if err != nil {
		log.Panic(err)
	}
	currCourse.Instructor = instructor

	days, err := getTextFieldsFromTokenizer(tokenizer, []string{"td"}, "days")
	if err != nil {
		log.Panic(err)
	}
	currCourse.Days = days

	//Special case where begin/end are merged into a single cell
	if days == "(ARR)" {
		time, err := getTextFieldsFromTokenizer(tokenizer, []string{"td"}, "time")
		if err != nil {
			log.Panic(err)
		}
		currCourse.TimeBegin = time
		currCourse.TimeEnd = time
	} else {
		timeBegin, err := getTextFieldsFromTokenizer(tokenizer, []string{"td"}, "timeBegin")
		if err != nil {
			log.Panic(err)
		}
		currCourse.TimeBegin = timeBegin

		timeEnd, err := getTextFieldsFromTokenizer(tokenizer, []string{"td"}, "timeEnd")
		if err != nil {
			log.Panic(err)
		}
		currCourse.TimeEnd = timeEnd
	}

	location, err := getTextFieldsFromTokenizer(tokenizer, []string{"td"}, "location")
	if err != nil {
		log.Panic(err)
	}
	currCourse.Location = location

	examTime, err := getTextFieldsFromTokenizer(tokenizer, []string{"td", "a"}, "examTime")
	if err != nil {
		log.Panic(err)
	}
	currCourse.ExamTime = examTime

	//Read next lines to check for additional times
	//Check if the table contains "Additional Times", this occurs 7 newline elements later in the table (multiple <td> elements)
	currentBuffer := string(tokenizer.Buffered())
	lines := strings.Split(currentBuffer, "\n")
	if len(lines) > 0 && strings.Contains(lines[7], "Additional Times") {
		//Do the case here, reading the row and setting the additional times field
		var at AdditionalTimes
		additionalTimesDays, err := getTextFieldsFromTokenizer(tokenizer, []string{"tr", "td", "td", "td", "td", "td", "b", "td"}, "additionalTimesDays")
		if err != nil {
			log.Panic(err)
		}
		at.Days = additionalTimesDays

		additionalTimeBegin, err := getTextFieldsFromTokenizer(tokenizer, []string{"td"}, "additionalTimesBegin")
		if err != nil {
			log.Panic(err)
		}
		at.TimeBegin = additionalTimeBegin

		additionalTimeEnd, err := getTextFieldsFromTokenizer(tokenizer, []string{"td"}, "additionalTimesEnd")
		if err != nil {
			log.Panic(err)
		}
		at.TimeEnd = additionalTimeEnd

		additionalLocation, err := getTextFieldsFromTokenizer(tokenizer, []string{"td"}, "additionalLocation")
		if err != nil {
			log.Panic(err)
		}
		at.Location = additionalLocation

		//Send tokenizer to end of line
		_, err = getTextFieldsFromTokenizer(tokenizer, []string{"td"}, "skipToRowEnd")
		if err != nil {
			log.Panic(err)
		}
		currCourse.AdditionalTimes = at
	}

	return currCourse, true
}
