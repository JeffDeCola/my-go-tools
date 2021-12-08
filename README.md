# my-go-tools

[![Go Report Card](https://goreportcard.com/badge/github.com/JeffDeCola/my-go-tools)](https://goreportcard.com/report/github.com/JeffDeCola/my-go-tools)
[![GoDoc](https://godoc.org/github.com/JeffDeCola/my-go-tools?status.svg)](https://godoc.org/github.com/JeffDeCola/my-go-tools)
[![Maintainability](https://api.codeclimate.com/v1/badges/ad5212958a91606b685e/maintainability)](https://codeclimate.com/github/JeffDeCola/my-go-tools/maintainability)
[![Issue Count](https://codeclimate.com/github/JeffDeCola/my-go-tools/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/my-go-tools/issues)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org)

_A place to keep useful tools I created in go._

Table of Contents,

* [CRYPTOGRAPHY TOOLS](https://github.com/JeffDeCola/my-go-tools#cryptography-tools)
* [MARKDOWN TOOLS](https://github.com/JeffDeCola/my-go-tools#markdown-tools)
* [CONTINUOUS INTEGRATION](https://github.com/JeffDeCola/my-go-tools#continuous-integration)

Documentation and reference,

* My [go-cheat-sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet)
* [my-go-examples](https://github.com/JeffDeCola/my-go-examples)
* [my-go-packages](https://github.com/JeffDeCola/my-go-packages)

Documentation and references,

* This repos
  [github webpage](https://jeffdecola.github.io/my-go-tools/)
  _built with
  [concourse](https://github.com/JeffDeCola/my-go-tools/blob/master/ci-README.md)_

## CRYPTOGRAPHY TOOLS

TESTSTSETSETSETSETSET

* [decryptfile](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile)

  _Decrypt a file with AES-256 GCM (a 32-byte hash key) using the `crypto/aes` package.
  Works with
  [encryptfile](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/encryptfile)._

* [encryptfile](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/encryptfile)

  _Encrypt a file with AES-256 GCM (a 32-byte hash key) using the `crypto/aes` package.
  Works with
  [decryptfile](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/decryptfile)._

* [md5-hash-file](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/md5-hash-file)

  _Get an md5 hash (fingerprint) from an input file using the standard
  `crypto/md5` package.
  I also added a flag to read in your `.ssh/id_rsa.pub` key to get your ssh fingerprint._

## MARKDOWN TOOLS

* [markdown-create-table-of-contents](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents)

  _Parse a markdown file to find `##`, `###` to create a table of contents (TOC)
  for links at github._

* [markdown-delimiter-doer](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer)

  _Take a markdown file and "do whatever you want" between the delimiters
  and output new markdown file._
  
## CONTINUOUS INTEGRATION

Refer to
[ci-README.md](https://github.com/JeffDeCola/my-go-tools/blob/master/ci-README.md)
on how I automated building this repos
[github webpage](https://jeffdecola.github.io/my-go-tools/)
and perform some unit-tests.
