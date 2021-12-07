package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const toolVersion = "2.0.1"

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

func readFile(filename string) []byte {

	log.Trace("Read the file ", filename, " to encrypt")
	fileDataToEncrypt, err := ioutil.ReadFile(filename)
	checkErr(err)
	log.Trace("Data/File to encrypt\n--------------------\n", string(fileDataToEncrypt), "\n--------------------\n")
	return fileDataToEncrypt

}

func getParaphrase() string {

	log.Trace("Get the paraphrase")
	paraphrase := ""
	fmt.Print("\nWhat is your secret paraphrase? ")
	_, err := fmt.Scan(&paraphrase)
	checkErr(err)
	return paraphrase
}

func getKeyByte(paraphrase string) []byte {

	log.Trace("hash the paraphrase'", paraphrase, "'to get 32 byte key")
	hasher := md5.New()
	hasher.Write([]byte(paraphrase))
	hash := hex.EncodeToString(hasher.Sum(nil))
	log.Trace("Hash is ", hash)

	keyByte := []byte(hash)
	log.Trace("keybyte is ", keyByte)
	return keyByte

}

func encryptFileData(keyByte []byte, plainTextByte []byte) string {

	// ENCRYPT DATA WITH 32 BYTE KEY AND RETURN CIPHERTEXT

	log.Trace("Encrypt file with key")
	fmt.Println("Encrypting input file")

	// GET CIPHER BLOCK USING KEY
	block, err := aes.NewCipher(keyByte)
	checkErr(err)

	// GET GCM INSTANCE THAT USES THE AES CIPHER
	gcm, err := cipher.NewGCM(block)
	checkErr(err)

	// CREATE A NONCE
	nonce := make([]byte, gcm.NonceSize())
	// Populates our nonce with a cryptographically secure random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// ENCRYPT DATA
	// Note how we put the Nonce in the beginging,
	// So we can rip it out when we decrypt
	cipherTextByte := gcm.Seal(nonce, nonce, plainTextByte, nil)

	// RETURN HEX
	cipherText := hex.EncodeToString(cipherTextByte)
	log.Trace("Encrypted Data\n--------------------\n", cipherText, "\n--------------------\n")
	return cipherText

}

func writeCipherText(cipherText string, filename string) {

	// Write cipherTex TO A FILE
	log.Trace("Write cipherText to a file")

	// Create file
	outputFile, err := os.Create(filename)
	checkErr(err)
	defer outputFile.Close()

	_, err = outputFile.WriteString("\nThis secret file was created by Jeff DeCola\n")
	checkErr(err)
	t := time.Now()
	_, err = outputFile.WriteString(t.Format(time.ANSIC) + "\n")
	checkErr(err)
	_, err = outputFile.WriteString("\n--------------------------------------------------------------------------------\n")
	checkErr(err)

	// Chop up the cipherText into lines of 80 characters into a slice.
	a := []rune(cipherText)
	line := ""
	numberLines := 0
	numberCharacters := 0

	for i, r := range a {
		line = line + string(r)
		if i > 0 && (i+1)%80 == 0 {
			line = line + "\n"
			_, err = outputFile.WriteString(line)
			checkErr(err)
			// Reset line
			line = ""
			numberLines++
		}
		numberCharacters++
	}

	if line != "" {
		// The remaining line
		numberLines++
		line = line + "\n"
		_, err = outputFile.WriteString(line)
		checkErr(err)
	}

	fmt.Printf("There were %v characters and %v lines created\n", numberCharacters, numberLines)

	_, err = outputFile.WriteString("--------------------------------------------------------------------------------\n\n")
	checkErr(err)

	fmt.Printf("Wrote output file\n\n")

}

func main() {

	// FLAGS
	version := flag.Bool("v", false, "prints current version")
	debugTrace := flag.Bool("debug", false, "log trace level")
	inputFilenamePtr := flag.String("i", "INPUT", "input file")
	outputFilenamePtr := flag.String("o", "OUTPUT", "output file")
	flag.Parse()

	log.Trace("inputFilename is ", *inputFilenamePtr)
	log.Trace("outputFilename is ", *outputFilenamePtr)

	// CHECK VERSION
	checkVersion(*version)

	// CHECK LOG LEVEL
	setLogLevel(*debugTrace)

	// GET DATA - Read the file - Will be a slice of bytes
	fileDataToEncrypt := readFile(*inputFilenamePtr)

	// GET PARAPHRASE - Ask the User
	paraphrase := getParaphrase()

	// GET KEY BYTE - Hash the paraphrase to get 32 Byte Key
	keyByte := getKeyByte(paraphrase)

	// ENCRYPT FILE DATA
	cipherText := encryptFileData(keyByte, fileDataToEncrypt)

	// WRITE cipherText TO FILE
	writeCipherText(cipherText, *outputFilenamePtr)

}
