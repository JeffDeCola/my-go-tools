# sha256-hash-file tool

A useful tool for
getting an sha256 hash (fingerprint) from an input file using the standard
`crypto/sha256` package.
I also added a flag to read in your .ssh/id_rsa.pub key to get your ssh sha256 fingerprint._

tl;dr,

```bash
# INSTALL VIA GO
go install sh256-hash-file.go

# GET SH256 FINGERPRINT ON FILE
sha256-hash-file testfile.txt

# GET SH256 FINGERPRINT ON KEY FILE
sha256-hash-file -ssh ~/.ssh/id_rsa.pub
```

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/sha256-hash-file#overview)
* [PREREQUISITES](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/sha256-hash-file#prerequisites)
* [RUN](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/sha256-hash-file#run)
* [TEST](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/sha256-hash-file#test)
* [INSTALL](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/sha256-hash-file#install)
* [USAGE](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/sha256-hash-file#usage)
  * [-h](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/sha256-hash-file#-h)
  * [-v](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/sha256-hash-file#-v)
  * [-ssh](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/sha256-hash-file#-ssh)
  * [-loglevel-string](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/sha256-hash-file#-loglevel-string)

Documentation and references,

* Refer to
  [sha256-hash-from-file](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/hashing/sha256-hash-from-file)
  example in `my-go-examples`
* This repos
  [github webpage](https://jeffdecola.github.io/my-go-tools/)
  _built with
  [concourse](https://github.com/JeffDeCola/my-go-tools/blob/master/ci-README.md)_

## OVERVIEW

In a nutshell, the guts of the code is,

```go
plainTextBytes := []byte(plainText)

// HASH
sha256HashByte := sha256.Sum256(plainTextBytes)

// CONVERT TO STRING
sha256Hash := base64.RawStdEncoding.EncodeToString(sha256HashByte[:])
```

## PREREQUISITES

I used the following language,

* [go](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet)

You will need the following go packages,

```bash
go get -u -v github.com/sirupsen/logrus
```

## RUN

The following steps are located in
[run.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/cryptography-tools/sha256-hash-file/run.sh).

To run
[sha256-hash-file.go](https://github.com/JeffDeCola/my-go-tools/blob/master/cryptography-tools/sha256-hash-file/sha256-hash-file.go)
from the command line,

```bash
go run . testfile.txt
go run sha256-hash-file.go testfile.txt
go run sha256-hash-file.go testfile.txt -loglevel trace
```

If you run on testfile.txt your sha256 hash shall be,

```txt
r2pN5Sjq43jggvxvW34cT6s662//HfyemRss38ZA+vk
```

You may also use the -ssh flag to read in a public ssh key file,

```bash
go run . -ssh ~/.ssh/id_rsa.pub
go run sha256-hash-file.go -ssh ~/.ssh/id_rsa.pub
go run sha256-hash-file.go -ssh ~/.ssh/id_rsa.pub -loglevel trace
```

Its nice because you can check you have the right fingerprint at
[github.com/settings/keys](https://github.com/settings/keys).

You can also run the unix command to check your sha256,

```bash
 ssh-keygen -lf ~/.ssh/id_rsa.pub
```

## TEST

The following steps are located in
[unit-tests.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/cryptography-tools/sha256-hash-file/test/unit-tests.sh).

To create `_test` files,

```bash
gotests -w -all sha256-hash-file.go
```

To unit test the code,

```bash
go test -cover ./... | tee test/test_coverage.txt
cat test/test_coverage.txt
```

## INSTALL

Will place an executable in your go bin,

```bash
go install sha256-hash-file.go
```

## USAGE

```txt
sha256-hash-file {-h|-v|-ssh} [FILENAME] -loglevel [level]
```

### -h

Help,

```bash
sha256-hash-file -h
```

### -v

Version,

```bash
sha256-hash-file -v
```

### -ssh

Check your public ssh file,

```bash
sha256-hash-file -ssh ~/.ssh/id_rsa.pub
```

### -loglevel string

Can use trace, info or error,

```bash
sha256-hash-file -ssh ~/.ssh/id_rsa.pub -loglevel trace
```
