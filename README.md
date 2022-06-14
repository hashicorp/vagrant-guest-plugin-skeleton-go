# vagrant-guest-plugin-skeleton

Ever wished you could spin up a Vagrant guest plugin right quick? Now, with vagrant-guest-plugin-skeleton you can! Follow these steps:

1. Copy this repo. Its a template so follow the Github instructions to [create a repo from a template using the UI](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-repository-from-a-template#creating-a-repository-from-a-template), or use the github cli.
```
$ gh repo create vagrant-guest-myspecialguest --template soapy1/vagrant-guest-plugin-skeleton --public
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
