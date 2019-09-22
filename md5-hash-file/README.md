# md5-hash-file tool

`md5-hash-file` _is a useful tool for
getting an md5 hash (fingerprint) from an input file using the standard
`crypto/md5` package.
I also added a flag to read in your `.ssh/id_rsa.pub` key to get your ssh fingerprint.

I expanded
[md5-hash-from-file](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/hashing/md5-hash-from-file)
example in `my-go-examples` to include a way to get the
md5 fingerprint from a ssh key.

[GitHub Webpage](https://jeffdecola.github.io/my-go-tools/)

## RUN

```bash
go run md5-hash-file.go -v
go run md5-hash-file.go <FILENAME>
go run md5-hash-file.go test.txt
```

If you run on test.txt your md5 hash shall be,

```txt
950dc9055bc2eb9b1f143e92d7bee6c4
```

You may also use the -ssh flag to read in a public ssh key file,

```bash
go run md5-hash-file.go -ssh <FILENAME>
go run md5-hash-file.go -ssh ~/.ssh/id_rsa.pub
go run md5-hash-file.go -ssh $HOME/.ssh/id_rsa.pub
```

Its nice because you can check you have the right fingerprint at
[github.com/settings/keys](https://github.com/settings/keys).

## INSTALL

```bash
go install md5-hash-file.go
```
