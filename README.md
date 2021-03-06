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

Documentation and reference,

* [go-cheat-sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet)
* [my-go-examples](https://github.com/JeffDeCola/my-go-examples)
* [my-go-packages](https://github.com/JeffDeCola/my-go-packages)

[GitHub Webpage](https://jeffdecola.github.io/my-go-tools/)

## USEFUL TOOLS I CREATED

_All sections in alphabetical order._

### CRYPTOGRAPHY TOOLS

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

### MARKDOWN TOOLS

* [markdown-create-table-of-contents](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-create-table-of-contents)

  _Parse a markdown file to find `##`, `###` to create a table of contents (TOC)
  for links at github.  I like to use a TOC at the beginning of my README.md files._

* [markdown-delimiter-doer](https://github.com/JeffDeCola/my-go-tools/tree/master/markdown-tools/markdown-delimiter-doer)

  _Take a markdown file and "do whatever you want" between the delimiters
  and output new markdown file._
  
## UPDATE GITHUB WEBPAGE USING CONCOURSE (OPTIONAL)

For fun, I use concourse to update
[my-go-tools GitHub Webpage](https://jeffdecola.github.io/my-go-tools/)
and alert me of the changes via repo status and slack.

A pipeline file [pipeline.yml](https://github.com/JeffDeCola/my-go-tools/tree/master/ci/pipeline.yml)
shows the entire ci flow. Visually, it looks like,

![IMAGE - my-go-tools concourse ci pipeline - IMAGE](docs/pics/my-go-tools-pipeline.jpg)

The `jobs` and `tasks` are,

* `job-readme-github-pages` runs task
  [readme-github-pages.sh](https://github.com/JeffDeCola/my-go-tools/tree/master/ci/scripts/readme-github-pages.sh).

The concourse `resources types` are,

* `my-go-tools` uses a resource type
  [docker-image](https://hub.docker.com/r/concourse/git-resource/)
  to PULL a repo from github.
* `resource-slack-alert` uses a resource type
  [docker image](https://hub.docker.com/r/cfcommunity/slack-notification-resource)
  that will notify slack on your progress.
* `resource-repo-status` uses a resource type
  [docker image](https://hub.docker.com/r/dpb587/github-status-resource)
  that will update your git status for that particular commit.

For more information on using concourse for continuous integration,
refer to my cheat sheet on [concourse](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet).
