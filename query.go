package main

import (
	"net/http"
	"reflect"
	"strings"
)

/*timetableURL constant is where we send our request*/
const timetableURL = "https://banweb.banner.vt.edu/ssb/prod/HZSKVTSC.P_ProcRequest"

/*CourseQuery is the query object sent directly to the VT Timetable server */
type CourseQuery struct {
	Campus       string `json:"CAMPUS"`
	TermYear     string `json:"TERMYEAR"`
	CoreCode     string `json:"CORE_CODE"`
	SubjectCode  string `json:"subj_code"`
	ScheduleType string `json:"SCHDTYPE"`
	CourseNumber string `json:"CRSE_NUMBER"`
	Crn          string `json:"crn"`
	OpenOnly     string `json:"open_only"`
}

/*SendQuery sends the query to the given URL.*/
func SendQuery(inputURL string, cq CourseQuery) (*http.Response, error) {
	//Required for request to process correctly
	if len(cq.CoreCode) == 0 {
		cq.CoreCode = "AR%"
	}

	req, err := http.NewRequest("POST", inputURL, nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	v := reflect.ValueOf(cq)
	var jsonTag string
	var tagOnly string
	var value string
	for i := 0; i < v.NumField(); i++ {
		jsonTag = string(reflect.TypeOf(cq).Field(i).Tag)         //Gets struct tag using reflection
		tagOnly = strings.Trim(string([]rune(jsonTag)[5:]), "\"") //Cuts off first 5 chars
		value = v.Field(i).Interface().(string)
		if len(value) > 0 {
			q.Add(tagOnly, value)
		}
	}
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)

	return resp, err
}

/*CreateEmptyQuery creates an empty query object which can be manipulated with setter methods in this file. */
func CreateEmptyQuery() CourseQuery {
	cq := CourseQuery{}
	return cq
}

/*SetCampus returns the caller CourseQuery object with the campus field set.*/
func (cq CourseQuery) SetCampus(Campus string) CourseQuery {
	cq.Campus = Campus
	return cq
}

/*SetTermYear returns the caller CourseQuery object with the term year field set.*/
func (cq CourseQuery) SetTermYear(TermYear string) CourseQuery {
	cq.TermYear = TermYear
	return cq
}

/*SetCoreCode returns the caller CourseQuery object with the term year field set.*/
func (cq CourseQuery) SetCoreCode(CoreCode string) CourseQuery {
	cq.CoreCode = CoreCode
	return cq
}

/*SetSubjectCode returns the caller CourseQuery object with the term year field set.*/
func (cq CourseQuery) SetSubjectCode(SubjectCode string) CourseQuery {
	cq.SubjectCode = SubjectCode
	return cq
}

/*SetScheduleType returns the caller CourseQuery object with the term year field set.*/
func (cq CourseQuery) SetScheduleType(ScheduleType string) CourseQuery {
	cq.ScheduleType = ScheduleType
	return cq
}

/*SetCourseNumber returns the caller CourseQuery object with the term year field set.*/
func (cq CourseQuery) SetCourseNumber(CourseNumber string) CourseQuery {
	cq.CourseNumber = CourseNumber
	return cq
}

/*SetCrn returns the caller CourseQuery object with the term year field set.*/
func (cq CourseQuery) SetCrn(Crn string) CourseQuery {
	cq.Crn = Crn
	return cq
}

/*SetOpenOnly returns the caller CourseQuery object with the term year field set.*/
func (cq CourseQuery) SetOpenOnly(OpenOnly string) CourseQuery {
	cq.OpenOnly = OpenOnly
	return cq
}
