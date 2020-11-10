package ics

import (
	"fmt"
	"io/ioutil"
	"time"
)

// Create an new isc File with event an time, email & description
func Create(filepath, start, end, summary, description string) {

	// Filepath and filename
	path := fmt.Sprintf("%s.ics", filepath)

	// Last modified
	modified := time.Now()

	// Params for file
	params := fmt.Sprintf("BEGIN:VCALENDAR\nVERSION:2.0\nPRODID:-//Apple Inc.//Mac OS X 10.15.7//EN\nCALSCALE:GREGORIAN\nBEGIN:VTIMEZONE\nTZID:Europe/Berlin\nBEGIN:DAYLIGHT\nTZOFFSETFROM:+0100\nRRULE:FREQ=YEARLY;BYMONTH=3;BYDAY=-1SU\nDTSTART:19810329T020000\nTZNAME:CEST\nTZOFFSETTO:+0200\nEND:DAYLIGHT\nBEGIN:STANDARD\nTZOFFSETFROM:+0200\nRRULE:FREQ=YEARLY;BYMONTH=10;BYDAY=-1SU\nDTSTART:19961027T030000\nTZNAME:CET\nTZOFFSETTO:+0100\nEND:STANDARD\nEND:VTIMEZONE\nBEGIN:VEVENT\nCREATED:20201110T095802Z\nUID:A17D8B07-832B-4774-836E-4A8D46088F01\nDTEND;TZID=Europe/Berlin:20201111T100000\nTRANSP:OPAQUE\nX-APPLE-TRAVEL-ADVISORY-BEHAVIOR:AUTOMATIC\nSUMMARY:%s\nLAST-MODIFIED:%s\nDTSTAMP:20201110T095915Z\nDTSTART;TZID=Europe/Berlin:20201111T090000\nSEQUENCE:1\nDESCRIPTION:%s\nBEGIN:VALARM\nX-WR-ALARMUID:D2ABDF1A-CEC2-4E42-BCFE-1D227CFF5E3F\nUID:D2ABDF1A-CEC2-4E42-BCFE-1D227CFF5E3F\nTRIGGER:-P1D\nX-APPLE-DEFAULT-ALARM:TRUE\nATTACH;VALUE=URI:Chord\nACTION:AUDIO\nEND:VALARM\nEND:VEVENT\nEND:VCALENDAR", summary, modified.Format("20060201T150405Z"), description)

	// Data for the file
	data := []byte(params)

	// Write new file
	err := ioutil.WriteFile(path, data, 0664)
	if err != nil {
		fmt.Println("Connot create vcs file: ", err)
	}

}