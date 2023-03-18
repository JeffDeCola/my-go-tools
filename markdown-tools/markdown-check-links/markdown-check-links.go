package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

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

func encryptPlainText(r io.Reader, keyByte []byte, plainTextByte []byte) (string, error) {

	// ENCRYPT PLAINTEXT WITH 32 BYTE KEY AND RETURN CIPHERTEXT
	log.Trace("Encrypt plainText with key")

	// GET CIPHER BLOCK USING KEY
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", fmt.Errorf("unable to get cipher block using keyByte: %w", err)
	}

	// GET GCM INSTANCE THAT USES THE AES CIPHER
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("unable to get gcm instance that uses the aes cipher: %w", err)
	}

	// CREATE A NONCE AND POPULATE
	nonce := make([]byte, gcm.NonceSize())
	// Populates our nonce with a cryptographically secure random sequence
	_, err = io.ReadFull(r, nonce)
	if err != nil {
		return "", fmt.Errorf("unable to populates our nonce with a random sequence: %w", err)
	}
	log.Info("Nonce is ", nonce)

	// ENCRYPT DATA
	// Note how we put the Nonce in the beginging,
	// So we can rip it out when we decrypt
	cipherTextByte := gcm.Seal(nonce, nonce, plainTextByte, nil)

	// RETURN HEX
	cipherText := hex.EncodeToString(cipherTextByte)
	log.Trace("Encrypted Data (cipherText)\n--------------------\n", cipherText, "\n--------------------\n")

	return cipherText, nil

}

func writeCipherTextFile(cipherText string, filename string) error {

	// WRITE cipherTex TO A FILE
	log.Trace("Write cipherText to a file")

	// CREATE AND OPEN FILE
	outputFH, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("unable to create file: %w", err)
	}
	defer outputFH.Close()

	// STEP 1 - WRITE HEADER
	err = writeHeader(outputFH)
	if err != nil {
		return fmt.Errorf("unable to create header for cipherText File: %w", err)
	}

	// STEP 2 - WRITE cipherText
	// Chop up into 80 character lines and write
	err = writeCipherText(cipherText, outputFH)
	if err != nil {
		return fmt.Errorf("unable to create cipherText lines for cipherText File: %w", err)
	}

	// STEP 3 - WRITE FOOTER
	err = writeFooter(outputFH)
	if err != nil {
		return fmt.Errorf("unable to create footer for cipherText File: %w", err)
	}

	log.Trace("Wrote cipherText to file\n\n")

	return nil

}

func writeHeader(handle io.Writer) error {

	// WRITE cipherTex Header TO A FILE
	log.Trace("Write cipherText Header to a file")

	_, err := fmt.Fprint(handle, "\nThis secret file was created by Jeff DeCola\n")
	if err != nil {
		return fmt.Errorf("unable to add text to file: %w", err)
	}
	t := time.Now()
	_, err = fmt.Fprint(handle, t.Format(time.ANSIC)+"\n")
	if err != nil {
		return fmt.Errorf("unable to add text to file: %w", err)
	}
	_, err = fmt.Fprint(handle, "\n"+myFileDelimiter+"\n")
	if err != nil {
		return fmt.Errorf("unable to add text to file: %w", err)
	}

	return nil
}

func writeCipherText(cipherText string, handle io.Writer) error {

	// WRITE cipherTex lines TO A FILE
	log.Trace("Write cipherText lines to a file")

	// CHOP UP cipherText INTO LINE OF 80 CHARACTERS INTO A SLICE.
	a := []rune(cipherText)
	line := ""
	numberLines := 0
	numberCharacters := 0

	for i, r := range a {
		line = line + string(r)
		if i > 0 && (i+1)%80 == 0 {
			line = line + "\n"
			_, err := fmt.Fprint(handle, line)
			if err != nil {
				return fmt.Errorf("unable to write chopped cipherText to file: %w", err)
			}
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
		_, err := fmt.Fprint(handle, line)
		if err != nil {
			return fmt.Errorf("unable to write to file: %w", err)
		}
	}

	log.Info("There were ", numberCharacters, " characters and ", numberLines, " lines created")

	return nil

}

func writeFooter(handle io.Writer) error {

	// WRITE FOOTER TO A FILE
	log.Trace("Write Footer to io.Writer")

	// ADD FOOTER TO BOTTOM OF FILE
	_, err := fmt.Fprint(handle, myFileDelimiter+"\n\n")
	if err != nil {
		return fmt.Errorf("unable to write to io.Writer: %w", err)
	}

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

	// STEP 1 - GET plainText FROM FILE
	plainTextByte, err := readFile(*inputFilenamePtr)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	// STEP 2 - GET PARAPHRASE - Ask the user or use a file
	paraphrase, err := getParaphrase(os.Stdin, *paraphraseFilePtr)
	if err != nil {
		log.Fatalf("Error getting paraphrase: %s", err)
	}

	// STEP 3 - GET KEY BYTE - Hash the paraphrase to get 32 Byte Key
	keyByte := getKeyByte(paraphrase)

	// STEP 4 - ENCRYPT plainText BASED ON PARAPHRASE TO GET cipherText
	cipherText, err := encryptPlainText(rand.Reader, keyByte, plainTextByte)
	if err != nil {
		log.Fatalf("Error getting cipherText: %s", err)
	}

	// STEP 5 - WRITE cipherText TO FILE
	err = writeCipherTextFile(cipherText, *outputFilenamePtr)
	if err != nil {
		log.Fatalf("Error writing cipherText to file: %s", err)
	}

	fmt.Println(" ")

}
