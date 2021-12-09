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

const toolVersion = "2.0.3"

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

func getFilename(ssh bool) string {

	// GET FILE NAME FROM ARGS
	filenameSlice := flag.Args()

	if len(filenameSlice) != 1 {
		err := errors.New("only one file allowed")
		checkErr(err)
	}
	filename := filenameSlice[0] // Make it a string

	return filename

}

func readFile(filename string) []byte {

	log.Trace("Read the file ", filename, " to MD5")
	plainTextByte, err := ioutil.ReadFile(filename)
	checkErr(err)
	log.Trace("Data/File to encrypt\n--------------------\n", string(plainTextByte), "\n--------------------\n")
	return plainTextByte

}

func parseSSHFile(ssh bool, plainTextByte []byte) string {

	var plainText string

	// If ssh public key file, we must parse it
	if ssh {
		// Parse the file because the file looks like `ssh-rsa AAA...ABC comments`
		// Hence parts[1] is the key
		parts := strings.Fields(string(plainTextByte))
		if len(parts) < 2 {
			log.Fatal("bad parse")
		}
		log.Trace("The parsed ssh key is: \n", parts[1], "\n\n")
		plainText = parts[1]
	} else {
		plainText = string(plainTextByte)
	}

	return plainText

}

func calculateMD5Hash(plainText string, isSSH bool) string {

	var plainTextByte []byte

	// Needed for ssh keys
	if isSSH {
		plainTextByte, _ = base64.StdEncoding.DecodeString(plainText)
	} else {
		plainTextByte = []byte(plainText)
	}

	// HASH
	md5HashByte := md5.Sum(plainTextByte)

	// CONVERT TO STRING
	md5Hash := hex.EncodeToString(md5HashByte[:])

	fmt.Printf("The md5 hash is: \n%s \n", md5Hash)

	return md5Hash
}

func printReadableMD5(md5Hash string) {

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
	debugTracePtr := flag.Bool("debug", false, "log trace level")
	sshPtr := flag.Bool("ssh", false, "ssh input file")
	flag.Parse()

	// CHECK VERSION
	checkVersion(*versionPtr)

	// SET LOG LEVEL
	setLogLevel(*debugTracePtr)

	log.Trace("Version flag = ", *versionPtr)
	log.Trace("Debug flag = ", *debugTracePtr)
	log.Trace("sshPointer = ", *sshPtr)

	fmt.Println(" ")

	// GET FILENAME
	filename := getFilename(*sshPtr)

	// GET DATA TO FINGERPRINT - Read the file - Will be a slice of bytes
	plainTextByte := readFile(filename)

	// PARSE plainTextByte IF SSH FLAG USED
	plainText := parseSSHFile(*sshPtr, plainTextByte)

	// CALCULATE MD5 HASH FROM STRING
	md5Hash := calculateMD5Hash(plainText, *sshPtr)

	fmt.Println(" ")

	// PRINT MORE READABLE FORM
	printReadableMD5(md5Hash)

	fmt.Printf("\n\n")

}
