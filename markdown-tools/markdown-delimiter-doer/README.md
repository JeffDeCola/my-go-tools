# markdown-delimiter-doer tool

`markdown-delimiter-doer` _is a useful tool for
taking a markdown file and "do whatever you want" between the delimiters
and output new markdown file._

[GitHub Webpage](https://jeffdecola.github.io/my-go-tools/)

## RUN

Run using delimiters `$$` and the `-htmltable` switch,

```bash
markdown-delimiter-doer -v
go run markdown-delimiter-doer.go -htmltable -delimiter \$\$ -i input.md -o output.md
```

To install (place an executable in your $GOPATH/bin),

## INSTALL

```bash
go install markdown-delimiter-doer.go
```

## SWITCHES

You can make many different switches to do different things.

### HTML TABLE SWITCH

Here is an illustration using the `-htmltable` switch,

![IMAGE - markdown-delimiter-doer - IMAGE](../../docs/pics/markdown-delimiter-doer.jpg)

It will even check the dates and strikethrough them automatically.
