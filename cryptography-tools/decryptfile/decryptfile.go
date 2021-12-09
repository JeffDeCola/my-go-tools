package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

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

func getCipherText(filename string) string {

	cipherText := ""

	// DATA - Read the file - Will be a slice of bytes
	log.Trace("Read the inputFilename to decrypt")
	// Open input file
	inputFile, err := os.Open(filename)
	checkErr(err)
	defer inputFile.Close()

	// Start scanning the input file line by line
	scanner := bufio.NewScanner(inputFile) // Increment the token

	for scanner.Scan() {

		// Read a line in file
		line := scanner.Text()

		// If you find a delimiter, get all the lines in between and place in a table.
		if line == "--------------------------------------------------------------------------------" {

			// Stay in here until you find another delimiter
			for scanner.Scan() {

				// Read next line nad Start Building the long long string
				line = scanner.Text()

				// Exit and build table when you find another delimiter
				if line == "--------------------------------------------------------------------------------" {
					break
				}
				cipherText = cipherText + line
			}
		}
	}

	log.Trace("cipherText is:\n--------------------\n", cipherText, "\n--------------------\n")
	return cipherText

}

func getParaphrase(paraphraseFile string) string {

	var paraphrase string

	if paraphraseFile != "" {
		fmt.Println("Getting the paraphrase from the file", paraphraseFile)
		paraphrase = string(readFile(paraphraseFile))
	} else {
		log.Trace("Get the paraphrase from User")
		fmt.Print("What is your secret paraphrase? ")
		_, err := fmt.Scan(&paraphrase)
		checkErr(err)
	}
	return paraphrase

}

func readFile(filename string) []byte {

	log.Trace("Read the file ", filename)
	fileData, err := ioutil.ReadFile(filename)
	checkErr(err)
	log.Trace("File Data\n--------------------\n", string(fileData), "\n--------------------\n")
	return fileData

}

func getKeyByte(paraphrase string) []byte {

	log.Trace("Hash the paraphrase'", paraphrase, "'to get 32 byte key")
	hasher := md5.New()
	hasher.Write([]byte(paraphrase))
	hash := hex.EncodeToString(hasher.Sum(nil))
	log.Trace("Hash is ", hash)

	keyByte := []byte(hash)
	log.Trace("keyByte is ", keyByte)
	return keyByte

}

// HASH THE PARAPHRASE TO GET 32 BYTE KEY
func createKey(paraphrase string) (string, error) {
	log.Trace("hashing the paraphrase")
	hasher := md5.New()
	hasher.Write([]byte(paraphrase))
	hash := hex.EncodeToString(hasher.Sum(nil))
	log.Trace("32 BYTE hash paraphrase is ", hash)
	return hash, nil
}

// DECRYPT DATA WITH 32 BYTE KEY AND RETURN PLAINTEXT
func decryptCipherText(keyByte []byte, cipherText string) []byte {

	cipherTextByte, _ := hex.DecodeString(cipherText)

	// GET CIPHER BLOCK USING KEY
	block, err := aes.NewCipher(keyByte)
	checkErr(err)

	// GET GCM BLOCK
	gcm, err := cipher.NewGCM(block)
	checkErr(err)

	// EXTRACT NONCE FROM cipherTextByte
	// Because I put it there
	nonceSize := gcm.NonceSize()
	nonce, cipherTextByte := cipherTextByte[:nonceSize], cipherTextByte[nonceSize:]

	// DECRYPT DATA
	plainTextByte, err := gcm.Open(nil, nonce, cipherTextByte, nil)
	log.Trace("Decrypted Data - plainTextByte\n--------------------\n", string(plainTextByte), "\n--------------------\n")
	checkErr(err)

	// RETURN STRING
	return plainTextByte
}

func writePlainTextByte(plainTextByte []byte, filename string) {

	// Write cipherTex TO A FILE
	log.Trace("Write plainTextByte to a file")
	// WRITE TO FILE
	f, err := os.Create(filename)
	checkErr(err)
	defer f.Close()
	f.Write(plainTextByte)
	log.Trace("Wrote output file\n")

}

func main() {

	// FLAGS
	versionPtr := flag.Bool("v", false, "prints current version")
	debugTracePtr := flag.Bool("debug", false, "log trace level")
	inputFilenamePtr := flag.String("i", "INPUT", "input file")
	outputFilenamePtr := flag.String("o", "OUTPUT", "output file")
	paraphraseFilePtr := flag.String("paraphrasefile", "", "use a file as a paraphrase")
	flag.Parse()

	// CHECK VERSION
	checkVersion(*versionPtr)

	// SET LOG LEVEL
	setLogLevel(*debugTracePtr)

	log.Trace("Version flag = ", *versionPtr)
	log.Trace("Debug flag = ", *debugTracePtr)
	log.Trace("Input Filename = ", *inputFilenamePtr)
	log.Trace("Output Filename = ", *outputFilenamePtr)
	log.Trace("Paraphrase File = ", *paraphraseFilePtr)

	fmt.Println(" ")

	// GET CIPHERTEXT (in bytes) FROM INPUT FILE
	cipherText := getCipherText(*inputFilenamePtr)

	// GET PARAPHRASE - Ask the User
	paraphrase := getParaphrase(*paraphraseFilePtr)

	// GET KEY BYTE - Hash the paraphrase to get 32 Byte Key
	keyByte := getKeyByte(paraphrase)

	// DECRYPT cipherText BASED ON PARAPHRASE to get FILE DATA
	plainTextByte := decryptCipherText(keyByte, cipherText)

	// WRITE plainTextByte TO FILE
	writePlainTextByte(plainTextByte, *outputFilenamePtr)

	fmt.Println(" ")

}
