# decryptfile tool

`decryptfile` _is a useful tool for
decryptfile a file with AES-256 GCM (a 32-byte hash key) using the `crypto/aes` package.
Works with
[encryptfile](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/encryptfile)._

Table of Contents,

* [RUN](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile#run)
* [TEST](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile#test)
* [INSTALL](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile#install)
* [SWITCHES](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile#switches)
  * [-h](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile#-h)
  * [-v](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile#-v)
  * [-i string, -o string](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile#-i-string--o-string)
* [HOW IT WORKS](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile#how-it-works)
  * [STEP 1 - CREATE A HASH KEY](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile#step-1---create-a-hash-key)
  * [STEP 2 - DECRYPT FILE WITH 32 BYTE HASH KEY](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile#step-2---decrypt-file-with-32-byte-hash-key)

Documentation and references,

* Use my other tool
  [encryptfile](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/encryptfile)
  to encrypt
* Refer to my
  [aes-256-gcm](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/symmetric-cryptography/aes-256-gcm)
  example in `my-go-examples` repo on how I did the decryption.
* This repos
  [github webpage](https://jeffdecola.github.io/my-go-tools/)
  _built with
  [concourse](https://github.com/JeffDeCola/my-go-tools/blob/master/ci-README.md)_

## RUN

The following steps are located in
[run.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/cryptography-tools/decryptfile/run.sh).

To run
[decryptfile.go](https://github.com/JeffDeCola/my-go-tools/blob/master/cryptography-tools/decryptfile/decryptfile.go)
from the command line,

```bash
go run decryptfile.go -i input.txt -o output.txt
```

## TEST

The following steps are located in
[unit-tests.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/cryptography-tools/decryptfile/test/unit-tests.sh).

To unit test the code,

```bash
go test -cover ./... | tee test/test_coverage.txt
cat test/test_coverage.txt
```

To create `_test` files,

```bash
gotests -w -all decryptfile.go
```

## INSTALL

Will place an executable in your go bin,

```bash
go install decryptfile.go
```

## SWITCHES

The default is to use README.md and create a table of contents,

```bash
decryptfile
```

### -h

Help,

```bash
decryptfile -h
```

### -v

Get version,

```bash
decryptfile -v
```

### -i string, -o string

Use a specific input file and output file`,

```bash
decryptfile -i input.txt -o output.txt
```

## HOW IT WORKS

The Advanced Encryption Standard, or AES, is a symmetric
block cipher chosen by the U.S. government to protect classified
information and is implemented in software and hardware throughout
the world to encrypt sensitive data.

We're going to use AES-256 GCM encryption from the standard go
[crypto/aes](https://golang.org/pkg/crypto/aes/)
package.

### STEP 1 - CREATE A HASH KEY

First you need a 32 byte key (AES-256).  Instead of typing a 32
character in, lets make it simple by turning a simple paraphrase into a key.
We will use the standard go
[crypto/md5](https://golang.org/pkg/crypto/md5/)
package.

```go
hasher := md5.New()
hasher.Write([]byte(paraphrase))
hash := hex.EncodeToString(hasher.Sum(nil))
```

### STEP 2 - DECRYPT FILE WITH 32 BYTE HASH KEY

The encryption was done using AES-256 GCM from my example
[aes-256-gcm](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/symmetric-cryptography/aes-256-gcm)
Refer to that example for a complete description.

This illustration may help,

![IMAGE - decryptfile - IMAGE](../../docs/pics/decryptfile.jpg)
