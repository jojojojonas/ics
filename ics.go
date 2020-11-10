package ics

import (
	"fmt"
	"io/ioutil"
	"time"
)

type Options struct {
	Filepath    string
	StartDate   Dates
	EndDate     Dates
	Summary     string
	Description string
}

type Dates struct {
	Date string
	Time string
}

// Format the date
func date(options Dates) string {

	// Parse date
	date, err := time.Parse("01.02.2006", options.Date)
	if err != nil {
		fmt.Println("Cannot parse start date: ", err)
	}

	// Parse time
	time, err := time.Parse("15:04:05", options.Time)
	if err != nil {
		fmt.Println("Connot parse start time: ", err)
	}

	// Return the value
	return fmt.Sprintf("%s%s", date.Format("20060201"), time.Format("T150405"))

}

// Create an new isc File with event an time, email & description
func Create(options Options) (created bool, err error) {

	// Last modified
	modified := time.Now()

	// Start date
	start := date(Dates{options.StartDate.Date, options.StartDate.Time})
	end := date(Dates{options.EndDate.Date, options.EndDate.Time})

	// Params for file
	params := fmt.Sprintf("BEGIN:VCALENDAR\nVERSION:2.0\nPRODID:-//Apple Inc.//Mac OS X 10.15.7//EN\nCALSCALE:GREGORIAN\nBEGIN:VTIMEZONE\nTZID:Europe/Berlin\nBEGIN:DAYLIGHT\nTZOFFSETFROM:+0100\nRRULE:FREQ=YEARLY;BYMONTH=3;BYDAY=-1SU\nDTSTART:19810329T020000\nTZNAME:CEST\nTZOFFSETTO:+0200\nEND:DAYLIGHT\nBEGIN:STANDARD\nTZOFFSETFROM:+0200\nRRULE:FREQ=YEARLY;BYMONTH=10;BYDAY=-1SU\nDTSTART:19961027T030000\nTZNAME:CET\nTZOFFSETTO:+0100\nEND:STANDARD\nEND:VTIMEZONE\nBEGIN:VEVENT\nCREATED:%s\nUID:A17D8B07-832B-4774-836E-4A8D46088F01\nDTEND;TZID=Europe/Berlin:%s\nTRANSP:OPAQUE\nX-APPLE-TRAVEL-ADVISORY-BEHAVIOR:AUTOMATIC\nSUMMARY:%s\nLAST-MODIFIED:%s\nDTSTAMP:20201110T095915Z\nDTSTART;TZID=Europe/Berlin:%s\nSEQUENCE:1\nDESCRIPTION:%s\nBEGIN:VALARM\nX-WR-ALARMUID:D2ABDF1A-CEC2-4E42-BCFE-1D227CFF5E3F\nUID:D2ABDF1A-CEC2-4E42-BCFE-1D227CFF5E3F\nTRIGGER:-P1D\nX-APPLE-DEFAULT-ALARM:TRUE\nATTACH;VALUE=URI:Chord\nACTION:AUDIO\nEND:VALARM\nEND:VEVENT\nEND:VCALENDAR", modified.Format("20060201T150405Z"), end, options.Summary, modified.Format("20060201T150405Z"), start, options.Description)

	// Data for the file
	data := []byte(params)

	// Write new file
	err = ioutil.WriteFile(options.Filepath, data, 0664)
	if err != nil {
		fmt.Println("Cannot create vcs file: ", err)
	}

	// Return finished
	return true, nil

}
