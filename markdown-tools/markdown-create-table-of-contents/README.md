# markdown-create-table-of-contents tool

`markdown-create-table-of-contents` _is a useful tool for
parsing a markdown file to find ##, ### to create a table
of contents (TOC) for links at github._

tl;dr,

```bash
# INSTALL VIA GO
go install markdown-create-table-of-contents.go

# GET TABLE OF CONTENTS FROM README.md MARKDOWN FILE
markdown-create-table-of-contents
markdown-create-table-of-contents -h3

# GET TABLE OF CONTENTS FROM input.md MARKDOWN FILE
markdown-create-table-of-contents -i input.md
```

Table of Contents,

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
  * [-debug](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#-debug)

Documentation and references,

* This repos
  [github webpage](https://jeffdecola.github.io/my-go-tools/)
  _built with
  [concourse](https://github.com/JeffDeCola/my-go-tools/blob/master/ci-README.md)_

## OVERVIEW

This tool is useful for creating a Table of Contents for .md files.

You will need to update the go code for your directory structure.

## PREREQUISITES

I used the following language,

* [go](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet)

You will need the following go packages,

```bash
go get -u -v github.com/sirupsen/logrus
```

## RUN

The following steps are located in
[run.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/markdown-tools/markdown-create-table-of-contents/run.sh).

To run
[markdown-create-table-of-contents.go](https://github.com/JeffDeCola/my-go-tools/blob/master/markdown-tools/markdown-create-table-of-contents/markdown-create-table-of-contents.go)
from the command line,

```bash
go run .
go run markdown-create-table-of-contents.go
go run markdown-create-table-of-contents.go -debug
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
markdown-create-table-of-contents {-h|-v|-debug} -i [FILENAME] -h3
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

Use a specific input file `input.md`,

```bash
markdown-create-table-of-contents -i input.md
```

### -h3

Include sub headings 3 `###`,

```bash
markdown-create-table-of-contents -i input.md -h3
```

### -debug

```bash
markdown-create-table-of-contents -i input.md -h3 -debug
```
