# sha256-hash-file TOOL

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Get an sha256 hash (fingerprint) from an input file using the
[crypto/sha256](https://golang.org/pkg/crypto/sha256/)
standard package.
I also added a flag to read in your `.ssh/id_rsa.pub` key to get your ssh fingerprint._

tl;dr

```bash
# INSTALL VIA GO
go install sh256-hash-file.go

# GET SH256 FINGERPRINT ON FILE
sha256-hash-file testfile.txt

# GET SH256 FINGERPRINT ON KEY FILE
sha256-hash-file -ssh ~/.ssh/id_rsa.pub
```

Table of Contents

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

Documentation and Reference

* [crypto/sha256](https://golang.org/pkg/crypto/sha256/)
  standard package
* Refer to
  [sha256-hash-from-file](https://github.com/JeffDeCola/my-go-examples#cryptography)
  example in
  [my-go-examples](https://github.com/JeffDeCola/my-go-examples)

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

You will need the following go packages,

```bash
go get -u -v github.com/sirupsen/logrus
```

## RUN

To
[run.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/cryptography-tools/sha256-hash-file/run.sh),

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

To create _test files,

```bash
gotests -w -all sha256-hash-file.go
```

To
[unit-tests.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/cryptography-tools/sha256-hash-file/test/unit-tests.sh),

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
