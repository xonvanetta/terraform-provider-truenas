# Terraform Provider Truenas

Truenas provider that will be used to create datasets,shares and more, also to manage truenas settings as services etc.

## TODO:

- [ ] Support NFS service config
    - [X] Basic implementation
    - [X] Baisc test
    - [ ] Missing `MountdPort`,`RpcstatdPort`,`RpclockdPort`
- [ ] Support Dataset/Volume Basic
    - [X] Basic resource implementation
    - [ ] Basic data implementation
    - [X] Baisc test
    - [ ] Permissions
    - [ ] Encryption
- [ ] Support Shares
    - [ ] NFS
        - [X] Basic implementation
        - [ ] Tests
    - [ ] SMB
        - [ ] Basic implementation
        - [ ] Tests
- [ ] Users
- [ ] Support Pools (Not knowing if it's a good idea or just rely on first time creation)
- [ ] Everything else from the api..

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 2.x?
- [Go](https://golang.org/doc/install) >= 1.16
- [Truenas](https://www.truenas.com/download-truenas-scale/) v2 api (should be installed on vm/bare metal machine)

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```sh
$ go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules). Please see the Go documentation for the most
up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```sh
$kgo get github.com/author/dependency
$ go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

See the [examples](examples) or check the docs

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (
see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin`
directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```