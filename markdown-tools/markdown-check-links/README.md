# markdown-check-links tool

A useful tool for checking links in a markdown file._

tl;dr,

```bash
# INSTALL VIA GO
go install markdown-check-links.go

# CHECK LINKS
markdown-check-links

# CHECK LINKS RECURSIVE
markdown-check-links -r
```

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/markdown-check-links#overview)
  * [STEP 1 - CREATE A HASH KEY](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/markdown-check-links#step-1---create-a-hash-key)
  * [STEP 2 - ENCRYPT FILE WITH 32 BYTE HASH KEY](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/markdown-check-links#step-2---encrypt-file-with-32-byte-hash-key)
* [PREREQUISITES](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/markdown-check-links#prerequisites)
* [RUN](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/markdown-check-links#run)
* [TEST](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/markdown-check-links#test)
* [INSTALL](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/markdown-check-links#install)
* [USAGE](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/markdown-check-links#usage)
  * [-h](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/markdown-check-links#-h)
  * [-v](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/markdown-check-links#-v)
  * [-i string](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/markdown-check-links#-i-string)
  * [-r](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/markdown-check-links#-r)
  * [-loglevel string](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile#-loglevel-string)

Documentation and references,

* This repos
  [github webpage](https://jeffdecola.github.io/my-go-tools/)
  _built with
  [concourse](https://github.com/JeffDeCola/my-go-tools/blob/master/ci-README.md)_

## OVERVIEW

Scan a markdown file and check the http links. Can also check subdirectories recursively.

## PREREQUISITES

I used the following language,

* [go](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet)

You will need the following go packages,

```bash
go install -v github.com/sirupsen/logrus
```

## RUN

The following steps are located in
[run.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/cryptography-tools/markdown-check-links/run.sh).

To run
[markdown-check-links.go](https://github.com/JeffDeCola/my-go-tools/blob/master/cryptography-tools/markdown-check-links/markdown-check-links.go)
from the command line,

```bash
go run .
go run markdown-check-links.go
go run markdown-check-links.go -i badlinks.md
go run markdown-check-links.go -i badlinks.md -loglevel trace
go run markdown-check-links.go -r
```

## TEST

The following steps are located in
[unit-tests.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/cryptography-tools/markdown-check-links/test/unit-tests.sh).

To create `_test` files,

```bash
gotests -w -all markdown-check-links.go
```

To unit test the code,

```bash
go test -cover ./... | tee test/test_coverage.txt
cat test/test_coverage.txt
```

## INSTALL

Will place an executable in your go bin,

```bash
go install markdown-check-links.go
```

## USAGE

```txt
markdown-check-links {-h|-v} {-i [input file]|-r} -loglevel [level]
```

### -h

Help,

```bash
markdown-check-links -h
```

### -v

Version,

```bash
markdown-check-links -v
```

### -i string

Use a specific input file. Will override `-r` flag,

```bash
markdown-check-links -i badlinks.md
```

### -r

Recursively check subdirectories. Will be ignored if `-i` flag is used,

```bash
markdown-check-links -r
```

### -loglevel string

Can be trace, info or error,

```bash
markdown-check-links -loglevel trace
```
