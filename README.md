# vagrant-guest-plugin-skeleton

Ever wished you could spin up a Vagrant guest plugin right quick? Now, with vagrant-guest-plugin-skeleton you can! Follow these steps:

1. Clone this repo
```
$ git clone soapy1/vagrant-guest-plugin-skeleton
```

2. Setup your go project
```
$ go mod init github.com/USERNAME/vagrant-guest-MYSPECIALGUEST
```

3. Update some bits
  - update the name of your guest in the `internal/guest` package and `main.go`
  - update the import statements in `main.go`

4. Tidy everything up
```
$ go mod tidy
$ go build .
```

5. Fill in the implementation!
