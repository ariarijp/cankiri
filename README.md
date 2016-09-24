cankiri
=====================

Show SSH Config as Tab Separated Values.

## Usage

```shell
$ go get github.com/ariarijp/cankiri
$ cankiri -h
Usage of cankiri:
  -format string
    	Format (default "tsv")
  -sep string
    	Separator (default "\t")
$ cankiri /path/to/ssh/config
Host	Host Name	User	Port	Identity File	SSH Command
foo	192.168.0.200	admin	2022		ssh -p 2022 admin@foo
bar	192.168.0.201	web	2222		ssh -p 2222 web@bar
```

## License

MIT

## Author

[Takuya Arita](https://github.com/ariarijp)
