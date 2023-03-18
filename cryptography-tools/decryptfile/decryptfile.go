package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

const toolVersion = "3.0.4"
const myFileDelimiter = "--------------------------------------------------------------------------------"

var errLogLevel = errors.New("please use trace, info or error")

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
		return fmt.Errorf("%s", errLogLevel)
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

func getCipherText(filename string) (string, error) {

	// get cipherText
	log.Trace("Get cipherText between delimiters")

	cipherText := ""

	// DATA - Open the file - Will be a slice of bytes
	log.Trace("Read the inputFilename to decrypt")
	// Open input file
	inputFile, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("unable to open file: %w", err)
	}
	defer inputFile.Close()

	// Start scanning the input file line by line
	scanner := bufio.NewScanner(inputFile) // Increment the token

	for scanner.Scan() {

		// Read a line in file
		line := scanner.Text()

		// If you find a delimiter, get all the lines in between and place in a table.
		if line == myFileDelimiter {

			// Stay in here until you find another delimiter
			for scanner.Scan() {

				// Read next line nad Start Building the long long string
				line = scanner.Text()

				// Exit and build table when you find another delimiter
				if line == myFileDelimiter {
					break
				}
				cipherText = cipherText + line
			}
		}
	}

	log.Trace("cipherText is:\n--------------------\n", cipherText, "\n--------------------\n")

	return cipherText, nil

}

func getParaphrase(r io.Reader, paraphraseFile string) (string, error) {

	var paraphrase string
	var err error

	// GET THE PARAPHRASE
	log.Trace("Get the paraphrase")

	// IS PARAPHRASE USER INPUT OR A FILE
	if paraphraseFile != "" {

		// FROM FILE
		log.Trace("Get the paraphrase from file")
		fmt.Println("Getting the paraphrase from the file", paraphraseFile)
		fileBytes, err := readFile(paraphraseFile)
		if err != nil {
			return "", fmt.Errorf("unable to open paraphrase: %w", err)
		}
		paraphrase = string(fileBytes)

	} else {

		// USER INPUT
		log.Trace("Get the paraphrase from User")
		paraphrase, err = getUserInput(r, "What is your secret paraphrase? ")
		if err != nil {
			return "", fmt.Errorf("unable to get get paraphrase: %w", err)
		}

	}

	return paraphrase, nil

}

func readFile(filename string) ([]byte, error) {

	// READ FILE
	log.Trace("Read the file ", filename)

	fileData, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %w", err)
	}
	log.Trace("File Data\n--------------------\n", string(fileData), "\n--------------------\n")

	return fileData, nil

}

func getUserInput(r io.Reader, askUser string) (string, error) {

	var nString string

	// GET STRING FROM USER
	log.Trace("Get string from user")
	fmt.Printf("%s", askUser)
	_, err := fmt.Fscan(r, &nString)
	if err != nil {
		return "", fmt.Errorf("unable to get string from user: %w", err)
	}

	return nString, nil

}

func getKeyByte(paraphrase string) []byte {

	// HASH THE PARAPHRASE
	log.Trace("Hash the paraphrase'", paraphrase, "'to get 32 byte key")

	// HASH USING MD5 HASH
	hasher := md5.New()
	hasher.Write([]byte(paraphrase))
	hash := hex.EncodeToString(hasher.Sum(nil))
	log.Info("Hashed paraphrase is ", hash)

	// MAKE KEYBYTE
	keyByte := []byte(hash)
	log.Info("Keybyte is ", keyByte)
	log.Trace("Keybyte string is ", string(keyByte))

	return keyByte

}

func decryptCipherText(keyByte []byte, cipherText string) ([]byte, error) {

	// DECRYPT DATA WITH 32 BYTE KEY AND RETURN PLAINTEXT
	log.Trace("Decrypt data with 32 byte key and return plaintext")

	cipherTextByte, _ := hex.DecodeString(cipherText)

	// GET CIPHER BLOCK USING KEY
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, fmt.Errorf("unable to get get cipher block: %w", err)
	}

	// GET GCM BLOCK
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("unable to get gcm block: %w", err)
	}

	// EXTRACT NONCE FROM cipherTextByte
	// Because I put it there
	nonceSize := gcm.NonceSize()
	nonce, cipherTextByte := cipherTextByte[:nonceSize], cipherTextByte[nonceSize:]

	// DECRYPT DATA
	plainTextByte, err := gcm.Open(nil, nonce, cipherTextByte, nil)
	log.Trace("Decrypted Data - plainTextByte\n--------------------\n", string(plainTextByte), "\n--------------------\n")
	if err != nil {
		return nil, fmt.Errorf("unable to decrypt cipherTextByte: %w", err)
	}

	// RETURN STRING
	return plainTextByte, nil
}

func writePlainText(plainTextByte []byte, filename string) error {

	// Write cipherTex TO A FILE
	log.Trace("Write plainTextByte to a file")

	// CREATE FILE
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("unable to create file: %w", err)
	}
	defer f.Close()

	// WRITE TO FILE
	f.Write(plainTextByte)
	log.Trace("Wrote output file\n")

	return nil

}

func main() {

	// FLAGS
	versionPtr := flag.Bool("v", false, "prints current version")
	logLevelPtr := flag.String("loglevel", "error", "log level (trace, info or error)")
	inputFilenamePtr := flag.String("i", "INPUT", "input file")
	outputFilenamePtr := flag.String("o", "OUTPUT", "output file")
	paraphraseFilePtr := flag.String("paraphrasefile", "", "use a file as a paraphrase")
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
	log.Trace("Input Filename = ", *inputFilenamePtr)
	log.Trace("Output Filename = ", *outputFilenamePtr)
	log.Trace("Paraphrase File = ", *paraphraseFilePtr)

	fmt.Println(" ")

	// STEP 1 - GET cipherText FROM FILE
	cipherText, err := getCipherText(*inputFilenamePtr)
	if err != nil {
		log.Fatalf("Error getting cipherText: %s", err)
	}

	// STEP 2 - GET PARAPHRASE - Ask the user or use a file
	paraphrase, err := getParaphrase(os.Stdin, *paraphraseFilePtr)
	if err != nil {
		log.Fatalf("Error getting paraphrase: %s", err)
	}

	// STEP 3 - GET KEY BYTE - Hash the paraphrase to get 32 Byte Key
	keyByte := getKeyByte(paraphrase)

	// STEP 4 - DECRYPT cipherText BASED ON KEYBYTE to get plainText
	plainTextByte, err := decryptCipherText(keyByte, cipherText)
	if err != nil {
		log.Fatalf("Error decrypting cipherText: %s", err)
	}

	// STEP 5 - WRITE plainTextByte TO FILE
	err = writePlainText(plainTextByte, *outputFilenamePtr)
	if err != nil {
		log.Fatalf("Error writing plainText to file: %s", err)
	}

	fmt.Println(" ")

}
