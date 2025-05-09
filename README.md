# MY GO TOOLS

[![Tag Latest](https://img.shields.io/github/v/tag/jeffdecola/my-go-tools)](https://github.com/JeffDeCola/my-go-tools/tags)
[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)
[![Go Reference](https://pkg.go.dev/badge/github.com/JeffDeCola/my-go-tools.svg)](https://pkg.go.dev/github.com/JeffDeCola/my-go-tools)
[![Go Report Card](https://goreportcard.com/badge/github.com/JeffDeCola/my-go-tools)](https://goreportcard.com/report/github.com/JeffDeCola/my-go-tools)

_A place to keep useful tools I created in go._

Table of Contents

* [CRYPTOGRAPHY TOOLS](https://github.com/JeffDeCola/my-go-tools#cryptography-tools)
* [MARKDOWN TOOLS](https://github.com/JeffDeCola/my-go-tools#markdown-tools)

Documentation and Reference

* [go-cheat-sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet)
* [my-go-examples](https://github.com/JeffDeCola/my-go-examples)
* [my-go-packages](https://github.com/JeffDeCola/my-go-packages)
* This repos
  [github webpage](https://jeffdecola.github.io/my-go-tools/)
  _built with
  [concourse](https://github.com/JeffDeCola/my-go-tools/blob/master/ci-README.md)_

## CRYPTOGRAPHY TOOLS

* [decryptfile](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile)

  _Decrypt a file with AES-256 GCM (a 32-byte hash key) using the
  [crypto/aes](https://golang.org/pkg/crypto/aes/)
  standard package.
  Works with
  [encryptfile](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/encryptfile)._

* [encryptfile](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/encryptfile)

  _Encrypt a file with AES-256 GCM (a 32-byte hash key) using the
  [crypto/aes](https://golang.org/pkg/crypto/aes/)
  standard package.
  Works with
  [decryptfile](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile)._

* [md5-hash-file](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/md5-hash-file)

  _Get an md5 hash (fingerprint) from an input file using the
  [crypto/md5](https://golang.org/pkg/crypto/md5/)
  standard package.
  I also added a flag to read in your `.ssh/id_rsa.pub` key to get your ssh fingerprint._

* [sha256-hash-file](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/sha256-hash-file)

  _Get an sha256 hash (fingerprint) from an input file using the
  [crypto/sha256](https://golang.org/pkg/crypto/sha256/)
  standard package.
  I also added a flag to read in your `.ssh/id_rsa.pub` key to get your ssh fingerprint._

## MARKDOWN TOOLS

* [markdown-check-links](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-check-links)

  _Check links in a markdown file._

* [markdown-create-table-of-contents](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents)

  _Parse a markdown file to find `##`, `###` to create a table of contents (TOC)
  for links at github._

* [markdown-delimiter-doer](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer)

  _Take a markdown file and "do whatever you want" between the delimiters
  and output new markdown file._
