# markdown-delimiter-doer TOOL

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Take a markdown file and "do whatever you want" between the delimiters
and output new markdown file._

tl;dr

```bash
# INSTALL VIA GO
go install markdown-delimiter-doer.go htmlswitch.go

# GET HTML FROM MARKDOWN
markdown-delimiter-doer -delimiter \$\$ -i input.md -o output.md -htmltable
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#overview)
* [PREREQUISITES](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#prerequisites)
* [RUN](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#run)
* [TEST](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#test)
* [INSTALL](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#install)
* [USAGE](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#usage)
  * [-h](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#-h)
  * [-v](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#-v)
  * [-delimiter string](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#-delimiter-string)
  * [-i string, -o string](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#-i-string--o-string)
  * [-htmltable](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#-htmltable)
  * [-loglevel string](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#-loglevel-string)
* [FUTURE](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#future)

## OVERVIEW

This tool is useful for "doing something" in a file between a delimiter.

## PREREQUISITES

You will need the following go packages,

```bash
go get -u -v github.com/sirupsen/logrus
```

## RUN

To
[run.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/markdown-tools/markdown-delimiter-doer/run.sh),

Run using delimiters `$$` and the `-htmltable` switch on input.md,

```bash
go run . -delimiter \$\$ \
       -i input.md -o output.md -htmltable
go run markdown-delimiter-doer.go htmlswitch.go \
       -delimiter \$\$ -i input.md -o output.md -htmltable
go run markdown-delimiter-doer.go htmlswitch.go \
       -delimiter \$\$ -i input.md -o output.md -htmltable -loglevel trace
```

## TEST

To create _test files,

```bash
gotests -w -all markdown-delimiter-doer.go
```

To
[unit-tests.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/markdown-tools/markdown-delimiter-doer/test/unit-tests.sh),

```bash
go test -cover ./... | tee test/test_coverage.txt
cat test/test_coverage.txt
```

## INSTALL

Will place an executable in your go bin,

```bash
go install markdown-delimiter-doer.go htmlswitch.go
```

## USAGE

```txt
markdown-delimiter-doer {-h|-v} -delimiter [delimiter]
                        -i [input file] -o [output file]
                        -htmltable -loglevel [level]
```

### -h

Help,

```bash
markdown-delimiter-doer -h
```

### -v

Version,

```bash
markdown-delimiter-doer -v
```

### -delimiter string

The delimiter you want to use.

### -i string, -o string

The `-i` and `-o` switches are used to define the input and output file respectively.

```bash
markdown-delimiter-doer -i input.md -o output.md
```

Running this command won't do anything, you need to use a switch like -htmltable.

### -htmltable

This switch is used to create an html table,

```bash
markdown-delimiter-doer -delimiter \$\$ -i input.md -o output.md -htmltable
```

Here is an illustration using the `-htmltable` switch,

![IMAGE - markdown-delimiter-doer - IMAGE](../../docs/pics/markdown-delimiter-doer.svg)

It will even check the dates and strikethrough them automatically.

### -loglevel string

Can be trace, info or error,

```bash
markdown-delimiter-doer -delimiter \$\$ -i input.md -o output.md -htmltable \
                        -loglevel trace
```

## FUTURE

The ability to add more switches to do whatever you want between delimiters.
