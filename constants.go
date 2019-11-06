package vttimetable

/*All of these "constants" are to be used when making API requests, as they are much more readable. */

/*CampusCodeConstants is a struct that would contain all of the constant values we need for campus codes. */
type CampusCodeConstants struct {
	BLACKSBURG            string
	VIRTUAL               string
	VTCSOM                string
	WESTERN               string
	VALLEY                string
	NATIONALCAPITALREGION string
	CENTRAL               string
	HAMPTONROADSCENTER    string
	CAPITAL               string
	OTHER                 string
}

/*CampusCodes is an aggregation of Campus Codes into a single object. Example: CampusCodes.BLACKSBURG returns "0" */
var CampusCodes = CampusCodeConstants{
	BLACKSBURG:            "0",
	VIRTUAL:               "10",
	VTCSOM:                "14",
	WESTERN:               "2",
	VALLEY:                "3",
	NATIONALCAPITALREGION: "4",
	CENTRAL:               "6",
	HAMPTONROADSCENTER:    "7",
	CAPITAL:               "8",
	OTHER:                 "9",
}

/*CoreCodeConstants is a type representing a Core Code (Core Curriculum) values aggregated into a struct to be only used in this file.*/
type CoreCodeConstants struct {
	ALLAREAS  string
	CLE1      string
	CLE2      string
	CLE3      string
	CLE4      string
	CLE5      string
	CLE6      string
	CLE7      string
	PATHWAY1A string
	PATHWAY1F string
	PATHWAY2  string
	PATHWAY3  string
	PATHWAY4  string
	PATHWAY5A string
	PATHWAY5F string
	PATHWAY6A string
	PATHWAY6D string
	PATHWAY7  string
}

/*CoreCodes is an aggregation of Core Codes into a single variable. Example: CoreCodes.ALLAREAS = "AR%" */
var CoreCodes = CoreCodeConstants{
	ALLAREAS:  "AR%",
	CLE1:      "AR01",
	CLE2:      "AR02",
	CLE3:      "AR03",
	CLE4:      "AR04",
	CLE5:      "AR05",
	CLE6:      "AR06",
	CLE7:      "AR07",
	PATHWAY1A: "G01A",
	PATHWAY1F: "G01F",
	PATHWAY2:  "G02",
	PATHWAY3:  "G03",
	PATHWAY4:  "G04",
	PATHWAY5A: "G05A",
	PATHWAY5F: "G05F",
	PATHWAY6A: "G06A",
	PATHWAY6D: "G06D",
	PATHWAY7:  "G07",
}

/*ScheduleTypeConstants is a type representing Schedule Type values aggregated into a single struct.*/
type ScheduleTypeConstants struct {
	ALLTYPES         string
	INDEPENDENTSTUDY string
	LAB              string
	LECTURE          string
	RECITATION       string
	RESEARCH         string
}

/*ScheduleType is an aggregation of Schedule Type constants into a single variable. Example: ScheduleType.ALLTYPES = "%" */
var ScheduleType = ScheduleTypeConstants{
	ALLTYPES:         "%",
	INDEPENDENTSTUDY: "%I%",
	LAB:              "%B%",
	LECTURE:          "%L%",
	RECITATION:       "%C%",
	RESEARCH:         "%R%",
}
