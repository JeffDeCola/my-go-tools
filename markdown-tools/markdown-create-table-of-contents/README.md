# markdown-create-table-of-contents tool

`markdown-create-table-of-contents` _is a useful tool for
parsing a markdown file to find ##, ### to create a table
of contents (TOC) for links at github._

Table of Contents,

* [RUN](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#run)
* [INSTALL](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#install)
* [HOW TO USE](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents#how-to-use)

Documentation and references,

* This repos
  [github webpage](https://jeffdecola.github.io/my-go-tools/)
  _built with
  [concourse](https://github.com/JeffDeCola/my-go-tools/blob/master/ci-README.md)_

## RUN

The following steps are located in run.sh,

```bash
go run markdown-create-table-of-contents.go
```

## INSTALL

Will place in your go bin,

```bash
go install markdown-create-table-of-contents.go
```

## HOW TO USE

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
markdown-create-table-of-contents.go -i input.md
```

Include sub headings 3 `###`,

```bash
markdown-create-table-of-contents.go -i input.md -h3
```
