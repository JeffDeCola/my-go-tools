package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const toolVersion = "2.0.3"

type table struct {
	columns       int
	colWdth       [20]string
	colAlgn       [20]string
	colBold       [20]string
	colDate       [20]string
	headers       [20]string
	rowColumnLine [40][20][10]string
	rows          int
}

func checkErr(err error) {

	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}

}

func checkVersion(version bool) {

	if version {
		fmt.Println(toolVersion)
		os.Exit(0)
	}

}

func setLogLevel(debugTrace bool) {

	// SET LOG LEVEL
	if debugTrace {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// SET FORMAT
	log.SetFormatter(&log.TextFormatter{})
	// log.SetFormatter(&log.JSONFormatter{})

	// SET OUTPUT (DEFAULT stderr)
	log.SetOutput(os.Stdout)

}

func doer(delimiter string, inputFilename string, outputFilename string, htmlTableBool bool) {

	// Temp storage for what you want to process between the delimiters
	var stuff []string

	// Open input file
	inputFile, err := os.Open(inputFilename)
	checkErr(err)
	defer inputFile.Close()

	// Create output file
	outputFile, err := os.Create(outputFilename)
	checkErr(err)

	// Start scanning the input file
	log.Trace("Start scanning the input file")
	scanner := bufio.NewScanner(inputFile) // Increment the token
	for scanner.Scan() {

		// Read a line in file
		line := scanner.Text()
		// fmt.Println("Working on:", line)

		// If you find a delimiter, get all the lines in between and place in a table.
		if line == delimiter {

			// Stay in here until you find another delimiter
			for scanner.Scan() {

				// Read a line in file
				line := scanner.Text()

				// Exit and build table when you find another delimiter
				if line == delimiter {
					break
				}

				// Place the line in stuff
				stuff = append(stuff, line)
			}

			// OK WE HAVE THE LINE ARRAY (Stuff between the delimiters)

			// htmltable switch
			if htmlTableBool {

				fmt.Println("MAKE HTML TABLE")
				makeHTMLTABLE(stuff, outputFile)
				fmt.Println("END MAKE HTML TABLE")
			}
			// reset stuff to empty
			stuff = []string{}

			// Did not find a delimiter on this line
		} else {

			// Write line to output file
			line := line + "\n"
			fmt.Print(line)
			_, err := outputFile.WriteString(line)
			if err != nil {
				fmt.Println(err)
				outputFile.Close()
				return
			}

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("ERROR:", err)
	}

}

func main() {

	// FLAGS
	versionPtr := flag.Bool("v", false, "prints current version")
	debugTracePtr := flag.Bool("debug", false, "log trace level")
	delimiterPtr := flag.String("delimiter", "DELIMETER", "what is the delimiter")
	inputFilenamePtr := flag.String("i", "INPUT", "input file")
	outputFilenamePtr := flag.String("o", "OUTPUT", "output file")
	htmlTableBoolPtr := flag.Bool("htmltable", false, "a bool")
	flag.Parse()

	// CHECK VERSION
	checkVersion(*versionPtr)

	// SET LOG LEVEL
	setLogLevel(*debugTracePtr)

	log.Trace("Version flag = ", *versionPtr)
	log.Trace("Debug flag = ", *debugTracePtr)
	log.Trace("delimiterPtr = ", *delimiterPtr)
	log.Trace("inputFilenamePtr = ", *inputFilenamePtr)
	log.Trace("outputFilenamePtr = ", *outputFilenamePtr)
	log.Trace("htmlTableBoolPtr = ", *htmlTableBoolPtr)

	fmt.Println(" ")

	doer(*delimiterPtr, *inputFilenamePtr, *outputFilenamePtr, *htmlTableBoolPtr)

	fmt.Println(" ")

}
