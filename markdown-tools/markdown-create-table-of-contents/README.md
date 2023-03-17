# markdown-create-table-of-contents tool

_Parse a markdown file to find `##`, `###` to create a table of contents (TOC)
for links at github._

tl;dr,

```bash
# INSTALL VIA GO
go install markdown-create-table-of-contents.go

# GET TABLE OF CONTENTS FROM README.md MARKDOWN FILE
# Will use tree/master in path
markdown-create-table-of-contents
markdown-create-table-of-contents -h3

# GET TABLE OF CONTENTS FROM input.md MARKDOWN FILE
# Will use blob/master in path
markdown-create-table-of-contents -i input.md
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#overview)
* [PREREQUISITES](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#prerequisites)
* [RUN](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#run)
* [TEST](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#test)
* [INSTALL](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#install)
* [USAGE](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#usage)
  * [-h](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#-h)
  * [-v](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#-v)
  * [-i string](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#-i-string)
  * [-h3](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#-h3)
  * [-loglevel string](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#-loglevel-string)

Documentation and Reference

* This repos
  [github webpage](https://jeffdecola.github.io/my-go-tools/)
  _built with
  [concourse](https://github.com/JeffDeCola/my-go-tools/blob/master/ci-README.md)_

## OVERVIEW

This tool is useful for creating a Table of Contents for .md files.

You will need to update the go code for your directory structure.

## PREREQUISITES

You will need the following go packages,

```bash
go get -u -v github.com/sirupsen/logrus
```

## RUN

To
[run.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/markdown-tools/markdown-create-table-of-contents/run.sh),

```bash
go run .
go run markdown-create-table-of-contents.go
go run markdown-create-table-of-contents.go -loglevel trace
```

## TEST

The following steps are located in
[unit-tests.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/markdown-tools/markdown-create-table-of-contents/test/unit-tests.sh).

To create `_test` files,

```bash
gotests -w -all markdown-create-table-of-contents.go
```

To unit test the code,

```bash
go test -cover ./... | tee test/test_coverage.txt
cat test/test_coverage.txt
```

## INSTALL

Will place an executable in your go bin,

```bash
go install markdown-create-table-of-contents.go
```

## USAGE

```txt
markdown-create-table-of-contents {-h|-v} -i [FILENAME] -h3 -loglevel [level]
```

The default is to use README.md and create a table of contents,

### -h

Help,

```bash
markdown-create-table-of-contents -h
```

### -v

Version,

```bash
markdown-create-table-of-contents -v
```

### -i string

Use a specific input file `input.md` which will use `blob/master` in path,

```bash
markdown-create-table-of-contents -i input.md
```

### -h3

Include sub headings 3 `###`,

```bash
markdown-create-table-of-contents -i input.md -h3
```

### -loglevel string

Can be trace, info or error,

```bash
markdown-create-table-of-contents -i input.md -h3 -loglevel trace
```
