# markdown-create-table-of-contents tool

`markdown-create-table-of-contents` _is a useful tool for
parsing a markdown file to find ##, ### to create a table
of contents (TOC) for links at github._

Table of Contents,

* [RUN](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#run)
* [TEST](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#test)
* [INSTALL](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#install)
* [USAGE](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#usage)

Documentation and references,

* This repos
  [github webpage](https://jeffdecola.github.io/my-go-tools/)
  _built with
  [concourse](https://github.com/JeffDeCola/my-go-tools/blob/master/ci-README.md)_

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

## USAGE

Default is to read README.md,

```bash
markdown-create-table-of-contents
```

Help,

```bash
markdown-create-table-of-contents -h
```

Get version,

```bash
markdown-create-table-of-contents -v
```

Use a specific filename `input.md`,

```bash
markdown-create-table-of-contents -i input.md
```

Include sub headings 3 `###`,

```bash
markdown-create-table-of-contents -i input.md -h3
```
