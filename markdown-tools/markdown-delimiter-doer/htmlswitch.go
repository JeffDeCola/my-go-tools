package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// Make an html table from the stuff
func makeHTMLTABLE(stuff []string, outputFile *os.File) {

	t := table{}

	buildTableStruct(&t, stuff)
	checkExpiredDates(&t)

	// <table>
	line := "<table style=\"font-size:.8em\">" + "\n"
	printLine(line, outputFile)

	// <col width> - Based on numbers of columns
	line = "  <!-- COLUMN WIDTHS -->" + "\n"
	printLine(line, outputFile)
	for i := 0; i < t.columns; i++ {
		line := "  <col width=\"" + t.colWdth[i] + "\">" + "\n"
		printLine(line, outputFile)
	}

	// <tr> - Heading row - Based on numbers of columns
	line = "<!-- HEADING ROW -->" + "\n"
	printLine(line, outputFile)
	line = "  <tr>" + "\n"
	printLine(line, outputFile)
	for i := 0; i < t.columns; i++ {
		line := "    <th>" + t.headers[i] + "</th>" + "\n"
		printLine(line, outputFile)
	}
	line = "  </tr>" + "\n"
	printLine(line, outputFile)

	// LETS DO THE ROWS - YEAH
	// ROWS - ITERATE OVER
	for r := 0; r < t.rows+1; r++ {
		line = "  <!-- ROW -->" + "\n"
		printLine(line, outputFile)
		line = "  <tr>" + "\n"
		printLine(line, outputFile)
		// COLUMNS - ITERATE OVER
		for c := 1; c < t.columns+1; c++ {
			alignment := t.colAlgn[c-1]
			bold := t.colBold[c-1]
			line = "    <td  align=\"" + alignment + "\" valign=\"top\">" + "\n"
			printLine(line, outputFile)
			// LINES - ITERATE OVER
			// Keep iterating until you get nothing
			for l := 0; l < 100; l++ {
				linebreak := ""
				prefixbold := ""
				suffixbold := ""
				// fmt.Println("r=", r, "c=", c, "l=", l)
				preline := t.rowColumnLine[r][c][l]
				// If blank do not print
				if preline != "" {
					// Add linebreak for multi lines
					if (l > 0) || (t.rowColumnLine[r][c][l+1] != "") {
						linebreak = "<br>"
					}
					// Add bold just on first line
					if (bold == "bold") && (l < 1) {
						prefixbold = "<b>"
						suffixbold = "</b>"
					}
					line = "      " + prefixbold + preline + suffixbold + linebreak + "\n"
					printLine(line, outputFile)
				} else {
					line = "    </td>" + "\n"
					printLine(line, outputFile)
					break
				}
			}
		}
		line = "  </tr>" + "\n"
		printLine(line, outputFile)
	}

	// <\table>
	line = "</table>" + "\n"
	printLine(line, outputFile)
}

// Place stuff in table struct
func buildTableStruct(t *table, stuff []string) {

	rowNumber := 0
	columnNumber := 0
	lineNumber := 0
	line := "stuff"

	// Loop over the array
	// Will get line by line
	for i := 0; i < len(stuff); i++ {

		line = stuff[i]

		if strings.Contains(line, "columns") {
			line = strings.Replace(line, "columns: ", "", -1)
			i, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
				return
			}
			t.columns = i
		}
		if strings.Contains(line, "colWdth") {
			line = strings.Replace(line, "colWdth: ", "", -1)
			foo := strings.Split(line, ",")
			for i, v := range foo {
				t.colWdth[i] = v
			}
		}
		if strings.Contains(line, "colAlgn") {
			line = strings.Replace(line, "colAlgn: ", "", -1)
			foo := strings.Split(line, ",")
			for i, v := range foo {
				t.colAlgn[i] = v
			}
		}
		if strings.Contains(line, "colBold") {
			line = strings.Replace(line, "colBold: ", "", -1)
			foo := strings.Split(line, ",")
			for i, v := range foo {
				t.colBold[i] = v
			}
		}
		if strings.Contains(line, "colDate") {
			line = strings.Replace(line, "colDate: ", "", -1)
			foo := strings.Split(line, ",")
			for i, v := range foo {
				t.colDate[i] = v
			}
		}
		if strings.Contains(line, "headers") {
			line = strings.Replace(line, "headers: ", "", -1)
			foo := strings.Split(line, ",")
			for i, v := range foo {
				t.headers[i] = v
			}
		}
		if strings.Contains(line, "rowcol") {
			// Get the column number
			tempcolumnNumberStr := string(line[6])
			tempcolumnNumber, _ := strconv.Atoi(tempcolumnNumberStr)
			if tempcolumnNumber < columnNumber {
				rowNumber++
				lineNumber = 0
			}
			if tempcolumnNumber == columnNumber {
				lineNumber++
			} else {
				lineNumber = 0
			}
			columnNumber = tempcolumnNumber
			replace := "rowcol" + tempcolumnNumberStr + ": "
			line := strings.Replace(line, replace, "", -1)
			log.Trace(replace, "rowNumber:", rowNumber, "columnNumber:", columnNumber, "lineNumber:", lineNumber, "----", line)
			// Place in 3-D array
			t.rowColumnLine[rowNumber][columnNumber][lineNumber] = line
		}

	}
	t.rows = rowNumber

}

// Go through all dates and strikethrew expired ones
func checkExpiredDates(t *table) {

	const layoutISO = "01/02/06"
	// GET CURRENT DATE
	// Format is MM/DD/YY
	currentTime := time.Now()
	// currentDate := currentTime.Format("01/02/06")

	// ROWS - ITERATE OVER
	for r := 0; r < t.rows+1; r++ {
		// COLUMNS - ITERATE OVER
		for c := 1; c < t.columns+1; c++ {
			// Do we check the date for this col
			if t.colDate[c-1] == "yes" {
				checkDate := t.rowColumnLine[r][c][0]
				// Place in go time format
				checkDateTime, _ := time.Parse(layoutISO, checkDate)
				if currentTime.After(checkDateTime) {
					t.rowColumnLine[r][c][0] = "<s>" + checkDate + "</s>"
				}

			}
		}
	}

}

// Print to output file
func printLine(line string, outputFile *os.File) {

	fmt.Print(line)
	_, err := outputFile.WriteString(line)
	if err != nil {
		fmt.Println(err)
		outputFile.Close()
		return
	}

}
