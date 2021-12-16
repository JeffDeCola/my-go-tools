// my-go-examples md5-hash-from-file.go

package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

const toolVersion = "2.0.4"

var ErrLogLevel = errors.New("please use trace, info or error")
var ErrFilenameArg = errors.New("only one filename allowed in args")
var ErrSSHFile = errors.New("can't parse ssh file correctly")

func setLogLevel(logLevel string) error {

	// SET LOG LEVEL (trace, info or error) None of which exit
	log.Trace("Set Log Level")

	switch logLevel {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.ErrorLevel)
		return fmt.Errorf("%s", ErrLogLevel)
	}

	// SET FORMAT
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})

	// SET OUTPUT (DEFAULT stderr)
	log.SetOutput(os.Stdout)

	return nil

}

func getFilename(ssh bool) (string, error) {

	// GET FILENAME FROM ARGS
	log.Trace("Get the filename from args")

	// GET FILE NAME FROM ARGS
	filenameSlice := flag.Args()

	if len(filenameSlice) != 1 {
		return "", fmt.Errorf("%s", ErrFilenameArg)
	}
	filename := filenameSlice[0] // Make it a string

	return filename, nil

}

func readFile(filename string) ([]byte, error) {

	// READ FILE
	log.Trace("Read the file ", filename)

	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %w", err)
	}
	log.Trace("File Data\n--------------------\n", string(fileData), "\n--------------------\n")

	return fileData, nil

}

func parseSSHFile(ssh bool, plainTextByte []byte) ([]byte, error) {

	var err error

	// PARSE SSH FILE DATA
	log.Trace("Parse the ssh data if flag set")

	// If ssh public key file, we must parse it
	if ssh {

		// Parse the file because the file looks like `ssh-rsa AAA...ABC comments`
		// Hence parts[1] is the key
		log.Trace("Check that you have a good ssh file")
		parts := strings.Fields(string(plainTextByte))
		if len(parts) < 2 {
			return nil, fmt.Errorf("%s", ErrSSHFile)
		}

		log.Trace("The parsed ssh key is: \n", parts[1], "\n\n")
		plainText := parts[1]
		plainTextByte, err = base64.StdEncoding.DecodeString(plainText)
		// plainTextByte = plainTextByteSSH
		if err != nil {
			return nil, fmt.Errorf("problem decoding string: %w", err)
		}

	} else {
		log.Trace("plainTextByte was just passed through")
	}

	return plainTextByte, nil

}

func calculateMD5Hash(plainTextByte []byte) string {

	// GET MD5 HASH from plainTextByte
	log.Trace("Get MD5 hash from plainTextByte")

	// HASH
	md5HashByte := md5.Sum(plainTextByte)

	// CONVERT TO STRING
	md5Hash := hex.EncodeToString(md5HashByte[:])

	log.Infof("The md5 hash is: \n%s \n", md5Hash)

	return md5Hash
}

func printReadableMD5(md5Hash string) {

	// PRINT MORE READABLE HASH
	log.Trace("Print a more readable Hash")

	// Get the hash in md5 bytes
	md5HashInBytes, _ := hex.DecodeString(md5Hash)

	// Print out the md5 fingerprint
	fmt.Println("The md5 fingerprint in a more readable form:")
	for i, b := range md5HashInBytes {
		fmt.Printf("%02x", b)
		if i < len(md5HashInBytes)-1 {
			fmt.Print(":")
		}
	}

}

func main() {

	// FLAGS
	versionPtr := flag.Bool("v", false, "prints current version")
	logLevelPtr := flag.String("loglevel", "error", "log level (trace, info or error)")
	sshPtr := flag.Bool("ssh", false, "ssh input file")
	flag.Parse()

	// CHECK VERSION
	if *versionPtr {
		fmt.Println(toolVersion)
		os.Exit(1)
	}

	// SET LOG LEVEL (trace, info or error) None of which exit
	err := setLogLevel(*logLevelPtr)
	if err != nil {
		log.Errorf("Error getting logLevel: %s", err)
	}

	// PRINT OUT FOR TRACE LOG
	log.Trace("Version flag = ", *versionPtr)
	log.Trace("Log Level = ", *logLevelPtr)
	log.Trace("using a ssh file = ", *sshPtr)

	fmt.Println(" ")

	// GET FILENAME FROM ARG
	filename, err := getFilename(*sshPtr)
	if err != nil {
		log.Fatalf("Error getting filename from arg: %s", err)
	}

	// GET DATA TO FINGERPRINT - Read the file - Will be a slice of bytes
	plainTextByte, err := readFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	// PARSE plainTextByte IF SSH FLAG USED
	plainTextByte, err = parseSSHFile(*sshPtr, plainTextByte)
	if err != nil {
		log.Fatalf("Can't Parse SHH FILE: %s", err)
	}

	// CALCULATE MD5 HASH FROM STRING
	md5Hash := calculateMD5Hash(plainTextByte)
	fmt.Printf("The md5 hash is: \n%s \n\n", md5Hash)

	// PRINT MORE READABLE FORM
	printReadableMD5(md5Hash)

	fmt.Println("")

}
