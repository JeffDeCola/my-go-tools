package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const toolVersion = "2.0.0"

var inputFilenamePtr, outputFilenamePtr *string

// HASH THE PARAPHRASE TO GET 32 BYTE KEY
func createKey(paraphrase string) (string, error) {
	log.Trace("hashing the paraphrase")
	hasher := md5.New()
	hasher.Write([]byte(paraphrase))
	hash := hex.EncodeToString(hasher.Sum(nil))
	log.Trace("hash is ", hash)
	return hash, nil
}

// DECRYPT DATA WITH 32 BYTE KEY AND RETURN PLAINTEXT
func decrypt(keyByte []byte, cipherText string) []byte {

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
	checkErr(err)

	// RETURN STRING
	return plainTextByte
}

func getCipherText(inputFile *os.File) string {

	cipherText := ""

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

	return cipherText

}

// Check your error
func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func init() {

	// SET LOG LEVEL
	// log.SetLevel(log.InfoLevel)
	log.SetLevel(log.TraceLevel)

	// SET FORMAT
	log.SetFormatter(&log.TextFormatter{})
	// log.SetFormatter(&log.JSONFormatter{})

	// SET OUTPUT (DEFAULT stderr)
	log.SetOutput(os.Stdout)

	// FLAGS
	version := flag.Bool("v", false, "prints current version")
	inputFilenamePtr = flag.String("i", "INPUT", "input file")
	outputFilenamePtr = flag.String("o", "OUTPUT", "output file")
	flag.Parse()

	// CHECK VERSION
	if *version {
		fmt.Println(toolVersion)
		os.Exit(0)
	}

}

func main() {

	log.Trace("inputFilename is %s, outputFilename is %s", *inputFilenamePtr, *outputFilenamePtr)

	// DATA
	// Read the file - Will be a slice of bytes
	log.Trace("Read the file to decrypt")
	// Open input file
	inputFile, err := os.Open(*inputFilenamePtr)
	checkErr(err)
	defer inputFile.Close()

	// GET CIPHERTEXT
	// (in bytes) FROM INPUT FILE
	cipherText := getCipherText(inputFile)
	// fmt.Printf("Data/File to decrypt\n--------------------\n%x\n--------------------\n", cipherText)
	fmt.Println("hi", cipherText, "bye")

	// PARAPHRASE
	// Ask the User
	log.Trace("Get the paraphrase")
	paraphrase := ""
	fmt.Printf("\nWhat is your secret paraphrase? ")
	_, err = fmt.Scan(&paraphrase)
	checkErr(err)

	// KEY
	// Has the paraphrase to get 32 Byte Key
	log.Trace("hash the paraphrase to get 32 byte key")
	keyText, err := createKey(paraphrase)
	keyByte := []byte(keyText)
	checkErr(err)

	// DECRYPT
	log.Trace("Decrypt cipherText with key")
	fmt.Println("Decrypting input file")
	plainTextByte := decrypt(keyByte, cipherText)
	// fmt.Printf("Decrypted Data\n--------------------\n%s\n--------------------\n", plainText)

	// WRITE TO FILE
	// Write plainText TO A FILE
	log.Trace("Write plainTextByte to a file")
	f, err := os.Create(*outputFilenamePtr)
	checkErr(err)
	defer f.Close()
	f.Write(plainTextByte)
	fmt.Printf("Wrote output file\n\n")

}
