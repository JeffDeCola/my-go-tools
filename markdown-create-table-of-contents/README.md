# markdown-create-table-of-contents tool

`markdown-create-table-of-contents` _is a useful tool for
parsing a markdown file to find ##, ### to create a table
of contents (TOC) for links at github.  I like to use a
TOC it at the beginning of my readme files._

For example, the following Table of Contents was generated using this tool,

* [RUN](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-create-table-of-contents#run)
* [INSTALL](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-create-table-of-contents#install)

[GitHub Webpage](https://jeffdecola.github.io/my-go-tools/)

## RUN

Default is README.md,

```bash
markdown-create-table-of-contents -v
go run markdown-create-table-of-contents
```

Use filename `input.md`,

```bash
go run markdown-create-table-of-contents.go -i input.md
```

Include headings 3 `###` switch `-h3`,

```bash
go run markdown-create-table-of-contents.go -i input.md -h3
```

## INSTALL

```bash
go install markdown-create-table-of-contents.go
```
