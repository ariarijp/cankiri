cankiri
=====================

Show SSH Config as Tab Separated Values or others.

## Usage

```shell
$ # Install
$ go get -u github.com/ariarijp/cankiri

$ # Usage
$ cankiri -h
Usage of cankiri:
  -format string
    	Format (default "tsv")
  -sep string
    	Separator (default "\t")

$ # Output as Tsv-Separated Values
$ go run main.go $GOPATH/src/github.com/ariarijp/cankiri/sample.config
Host	Host Name	User	Port	Identity File	Proxy Command	SSH Command
foo 192.168.0.200	admin	22			ssh -p 22 admin@192.168.0.200
bar 172.10.1.10	web	2222			ssh -p 2222 web@172.10.1.10

$ # Output as JSON
$ go run main.go -format json $GOPATH/src/github.com/ariarijp/cankiri/sample.config
[
  {"Host":"foo","HostName":"192.168.0.200","IdentityFile":"","Port":"22","ProxyCommand":"","SSHCmd":"ssh -p 22 admin@192.168.0.200","User":"admin"},
  {"Host":"bar","HostName":"172.10.1.10","IdentityFile":"","Port":"2222","ProxyCommand":"","SSHCmd":"ssh -p 2222 web@172.10.1.10","User":"web"}
]

$ # Output as JSONL
$ go run main.go -format jsonl $GOPATH/src/github.com/ariarijp/cankiri/sample.config
{"Host":"foo","HostName":"192.168.0.200","IdentityFile":"","Port":"22","ProxyCommand":"","SSHCmd":"ssh -p 22 admin@192.168.0.200","User":"admin"}
{"Host":"bar","HostName":"172.10.1.10","IdentityFile":"","Port":"2222","ProxyCommand":"","SSHCmd":"ssh -p 2222 web@172.10.1.10","User":"web"}
```

## License

MIT

## Author

[Takuya Arita](https://github.com/ariarijp)
