# markdown-delimiter-doer tool

`markdown-delimiter-doer` _is a useful tool for
taking a markdown file and "do whatever you want" between the delimiters
and output new markdown file._

Table of Contents,

* [RUN](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#run)
* [TEST](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#test)
* [INSTALL](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#install)
* [USAGE](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#usage)
  * [-i -o Switch](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#-i--o-switch)
  * [-htmltable Switch](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#-htmltable-switch)
* [FUTURE](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer#future)

Documentation and references,

* This repos
  [github webpage](https://jeffdecola.github.io/my-go-tools/)
  _built with
  [concourse](https://github.com/JeffDeCola/my-go-tools/blob/master/ci-README.md)_

## RUN

The following steps are located in
[run.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/markdown-tools/markdown-delimiter-doer/run.sh).

To run
[markdown-delimiter-doer.go](https://github.com/JeffDeCola/my-go-tools/blob/master/markdown-tools/markdown-delimiter-doer/markdown-delimiter-doer.go)
from the command line,

Run using delimiters `$$` and the `-htmltable` switch on input.md,

```bash
go run markdown-delimiter-doer.go -htmltable -delimiter \$\$ -i input.md -o output.md
```

## TEST

The following steps are located in
[unit-tests.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/markdown-tools/markdown-delimiter-doer/test/unit-tests.sh).

To unit test the code,

```bash
go test -cover ./... | tee test/test_coverage.txt
cat test/test_coverage.txt
```

To create `_test` files,

```bash
gotests -w -all markdown-delimiter-doer.go
```

## INSTALL

Will place an executable in your go bin,

```bash
go install markdown-delimiter-doer.go
```

## USAGE

Help,

```bash
markdown-delimiter-doer -h
```

Get version,

```bash
markdown-delimiter-doer -v
```

### -i -o Switch

The `-i` and `-o` switches are used to define the input and output file respectively.

```bash
markdown-delimiter-doer -i input.md -o output.md
```

### -htmltable Switch

This switch is used to create an html table,

```bash
markdown-delimiter-doer.go -htmltable -delimiter \$\$ -i input.md -o output.md
```

Here is an illustration using the `-htmltable` switch,

![IMAGE - markdown-delimiter-doer - IMAGE](../../docs/pics/markdown-delimiter-doer.jpg)

It will even check the dates and strikethrough them automatically.

## FUTURE

The ability to add more switches to do whatever you want between delimiters.
