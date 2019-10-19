package main

import (
	"log"
	"os"

	"golang.org/x/net/html"
)

/*Course represents a single course as read from the HTTP response that we are parsing.*/
type Course struct {
	Crn         string
	CourseName  string
	ClassName   string
	SectionType string
	CreditHours string
	Capacity    string
	Instructor  string
	Days        string
	TimeBegin   string
	TimeEnd     string
	Location    string
	ExamTime    string
}

func readFile(fileName string) []Course {
	file, err := os.Open(fileName)
	if err != nil {
		log.Panic("Error: could not open file")
	}
	tokenizer := html.NewTokenizer(file)

	courseList := make([]Course, 0)

	for {
		next := tokenizer.Next()

		switch {
		//html.ErrorToken can indicate EOF, or we return an error as well
		case next == html.ErrorToken:
			return courseList
		case next == html.StartTagToken:
			token := tokenizer.Token()

			for _, a := range token.Attr {
				if a.Key == "class" && a.Val == "dataentrytable" {
					//Start parsing through this table, adding to this object
					currCourse := Course{}

					//This is strictly for passing the first row of labels in the table
					err := getNextTR(tokenizer)
					if err != nil {
						log.Panic(err)
					}
					err = getNextTR(tokenizer)
					if err != nil {
						log.Panic(err)
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

					courseList = append(courseList, currCourse)
				}
			}
		}

	}

}

//Stops after reaching TR
func getNextTR(tokenizer *html.Tokenizer) error {
	for {
		next := tokenizer.Next()

		switch {
		case next == html.ErrorToken:
			return tokenizer.Err()
		case next == html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "tr" {
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
	return text.Data, err
}
