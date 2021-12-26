  _built with
  [concourse](https://github.com/JeffDeCola/my-go-tools/blob/master/ci-README.md)_

# CRYPTOGRAPHY TOOLS

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

* [sha256-hash-file](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/sha256-hash-file)

  _Get an sha256 hash (fingerprint) from an input file using the standard
  `crypto/sha256` package.
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
  
## CONTINUOUS INTEGRATION

Refer to
[ci-README.md](https://github.com/JeffDeCola/my-go-tools/blob/master/ci-README.md)
on how I automated building this repos
[github webpage](https://jeffdecola.github.io/my-go-tools/)
and perform some unit-tests.
