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

const toolVersion = "2.0.1"

var filename string

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func getFilename(filename string) (string, bool) {

	var isSSH bool

	// Check if the ssh flag is set, if not just use input file
	if filename == "" {
		isSSH = false
		// GET FILE NAME FROM ARGS
		filenameSlice := flag.Args()
		if len(filenameSlice) != 1 {
			err := errors.New("only one files allowed")
			checkErr(err)
		}
		filename = filenameSlice[0] // Make it a string

	} else {
		isSSH = true
	}

	return filename, isSSH
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

	return md5Hash
}

func init() {

	// FLAGS
	version := flag.Bool("v", false, "prints current version")
	debugTrace := flag.Bool("debug", false, "log trace level")
	filenamePtr := flag.String("ssh", "", "ssh input file")
	flag.Parse()
	filename = *filenamePtr

	// SET LOG LEVEL
	if *debugTrace {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// SET FORMAT
	log.SetFormatter(&log.TextFormatter{})
	// log.SetFormatter(&log.JSONFormatter{})

	// SET OUTPUT (DEFAULT stderr)
	log.SetOutput(os.Stdout)

	// CHECK VERSION
	if *version {
		fmt.Println(toolVersion)
		os.Exit(0)
	}

}

func main() {

	fmt.Println(" ")

	// GET FILENAME
	filename, isSSH := getFilename(filename)
	fmt.Printf("The filename is %s which is a public ssh key file: %v\n\n", filename, isSSH)

	// READ FILE INTO BYTES
	plainTextByte, err := ioutil.ReadFile(filename)
	checkErr(err)
	// Convert to string
	plainText := string(plainTextByte)
	fmt.Printf("The file %s contains: \n%s\n\n", filename, plainText)

	// If ssh public key file, we must parse it
	if isSSH {
		// Parse the file because the file looks like `ssh-rsa AAA...ABC comments`
		// Hence parts[1] is the key
		parts := strings.Fields(string(plainTextByte))
		if len(parts) < 2 {
			log.Fatal("bad parse")
		}
		fmt.Printf("The public ssh key is: \n%s \n\n", parts[1])
		plainText = parts[1]
	}

	// SO NOW WE ARE FINALLY READY
	// CALCULATE MD5 HASH FROM STRING
	md5Hash := calculateMD5Hash(plainText, isSSH)
	fmt.Printf("The md5 hash is: \n%s \n\n", md5Hash)

	if isSSH {
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

	fmt.Printf("\n\n")

}
