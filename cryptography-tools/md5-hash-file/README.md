# md5-hash-file TOOL

_Get an md5 hash (fingerprint) from an input file using the standard
`crypto/md5` package.
I also added a flag to read in your `.ssh/id_rsa.pub` key to get your ssh fingerprint._

tl;dr,

```bash
# INSTALL VIA GO
go install md5-hash-file.go

# GET MD5 FINGERPRINT ON FILE
md5-hash-file testfile.txt

# GET MD5 FINGERPRINT ON KEY FILE
md5-hash-file -ssh ~/.ssh/id_rsa.pub
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/md5-hash-file#overview)
* [PREREQUISITES](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/md5-hash-file#prerequisites)
* [RUN](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/md5-hash-file#run)
* [TEST](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/md5-hash-file#test)
* [INSTALL](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/md5-hash-file#install)
* [USAGE](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/md5-hash-file#usage)
  * [-h](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/md5-hash-file#-h)
  * [-v](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/md5-hash-file#-v)
  * [-ssh](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/md5-hash-file#-ssh)
  * [-loglevel string](https://github.com/JeffDeCola/my-go-tools/tree/master/cryptography-tools/md5-hash-file#-loglevel-string)

Documentation and Reference

* Refer to
  [md5-hash-from-file](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/hashing/md5-hash-from-file)
  example in
  [my-go-examples](https://github.com/JeffDeCola/my-go-examples)
* This repos
  [github webpage](https://jeffdecola.github.io/my-go-tools/)
  _built with
  [concourse](https://github.com/JeffDeCola/my-go-tools/blob/master/ci-README.md)_

## OVERVIEW

In a nutshell, the guts of the code is,

```go
plainTextBytes := []byte(plainText)

// HASH
md5HashByte := md5.Sum(plainTextBytes)

// CONVERT TO STRING
md5Hash := hex.EncodeToString(md5HashByte[:])
```

## PREREQUISITES

You will need the following go packages,

```bash
go get -u -v github.com/sirupsen/logrus
```

## RUN

To
[run.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/cryptography-tools/md5-hash-file/run.sh),

```bash
go run . testfile.txt
go run md5-hash-file.go testfile.txt
go run md5-hash-file.go testfile.txt -loglevel trace
```

If you run on testfile.txt your md5 hash shall be,

```txt
950dc9055bc2eb9b1f143e92d7bee6c4
```

You may also use the -ssh flag to read in a public ssh key file,

```bash
go run . -ssh ~/.ssh/id_rsa.pub
go run md5-hash-file.go -ssh ~/.ssh/id_rsa.pub
go run md5-hash-file.go -ssh ~/.ssh/id_rsa.pub -loglevel trace
```

Its nice because you can check you have the right fingerprint at
[github.com/settings/keys](https://github.com/settings/keys).

You can also run the unix command to check your md5,

```bash
ssh-keygen -l -E md5 -f ~/.ssh/id_rsa.pub
```

## TEST

To create _test files,

```bash
gotests -w -all md5-hash-file.go
```

To
[unit-tests.sh](https://github.com/JeffDeCola/my-go-tools/blob/master/cryptography-tools/md5-hash-file/test/unit-tests.sh),

```bash
go test -cover ./... | tee test/test_coverage.txt
cat test/test_coverage.txt
```

## INSTALL

Will place an executable in your go bin,

```bash
go install md5-hash-file.go
```

## USAGE

```txt
md5-hash-file {-h|-v|-ssh} [FILENAME] -loglevel [level] 
```

### -h

Help,

```bash
md5-hash-file -h
```

### -v

Version,

```bash
md5-hash-file -v
```

### -ssh

Check your public ssh file,

```bash
md5-hash-file -ssh ~/.ssh/id_rsa.pub
```

### -loglevel string

Can use trace, info or error,

```bash
md5-hash-file -ssh ~/.ssh/id_rsa.pub -loglevel trace
```
