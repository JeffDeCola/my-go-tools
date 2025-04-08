# markdown-check-links TOOL

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

```text
*** THIS IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

_Check links in a markdown file._

tl;dr

```bash
# INSTALL VIA GO
go install markdown-check-links.go

# CHECK LINKS
markdown-check-links

# CHECK LINKS RECURSIVE
markdown-check-links -r
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-check-links#overview)
* [PREREQUISITES](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-check-links#prerequisites)
* [RUN](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-check-links#run)
* [TEST](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-check-links#test)
* [INSTALL](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-check-links#install)
* [USAGE](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-check-links#usage)
  * [-h](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-check-links#-h)
  * [-v](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-check-links#-v)
  * [-i string](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-check-links#-i-string)
  * [-r](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-check-links#-r)
  * [-loglevel string](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-check-links#loglevel-string)

## OVERVIEW

Scan a markdown file and check the http links. Can also check subdirectories recursively.

## PREREQUISITES

You will need the following go packages,

```bash
go install -v github.com/sirupsen/logrus
```

## RUN

To
[run.sh](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-check-links/run.sh),

```bash
go run .
go run markdown-check-links.go
go run markdown-check-links.go -i badlinks.md
go run markdown-check-links.go -i badlinks.md -loglevel trace
go run markdown-check-links.go -r
```

## TEST

To create _test files,

```bash
gotests -w -all markdown-check-links.go
```

To
[unit-tests.sh](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-check-links/test/unit-tests.sh),

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
