# markdown-create-table-of-contents tool

`markdown-create-table-of-contents` _is a useful tool for
parsing a markdown file to find ##, ### to create a table
of contents (TOC) for links at github._

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#overview)
* [PREREQUISITES](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#prerequisites)
* [RUN](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#run)
* [TEST](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#test)
* [INSTALL](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#install)
* [SWITCHES](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#switches)
  * [-h](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#-h)
  * [-v](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#-v)
  * [-i string](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#-i-string)
  * [-h3](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#-h3)

Documentation and references,

* This repos
  [github webpage](https://jeffdecola.github.io/my-go-tools/)
  _built with
  [concourse](https://github.com/JeffDeCola/my-go-tools/blob/master/ci-README.md)_

## OVERVIEW

This tool is useful for creating a Table of Contents for .md files.

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
go run markdown-create-table-of-contents.go
```

## TEST

The following steps are located in
[unit-tests.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/markdown-tools/markdown-create-table-of-contents/test/unit-tests.sh).

To unit test the code,

```bash
go test -cover ./... | tee test/test_coverage.txt
cat test/test_coverage.txt
```

To create `_test` files,

```bash
gotests -w -all markdown-create-table-of-contents.go
```

## INSTALL

Will place an executable in your go bin,

```bash
go install markdown-create-table-of-contents.go
```

## SWITCHES

The default is to use README.md and create a table of contents,

```bash
markdown-create-table-of-contents
```

### -h

Help,

```bash
markdown-create-table-of-contents -h
```

### -v

Get version,

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
