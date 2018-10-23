# Library for INWX

[![Build Status](http://github.dronehippie.de/api/badges/webhippie/go-inwx/status.svg)](http://github.dronehippie.de/webhippie/go-inwx)
[![Stories in Ready](https://badge.waffle.io/webhippie/go-inwx.svg?label=ready&title=Ready)](http://waffle.io/webhippie/go-inwx)
[![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/78c149e77062448d9bd33efb38f0042e)](https://www.codacy.com/app/webhippie/go-inwx?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=webhippie/go-inwx&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/webhippie/go-inwx?status.svg)](http://godoc.org/github.com/webhippie/go-inwx)
[![Go Report](http://goreportcard.com/badge/github.com/webhippie/go-inwx)](http://goreportcard.com/report/github.com/webhippie/go-inwx)

This repository will provides helpers related to INWX.

## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.8. It is also possible to just simply execute the `go get github.com/webhippie/go-inwx/...` command, but we prefer to use our `Makefile`:

```bash
go get -d github.com/webhippie/go-inwx/...
cd $GOPATH/src/github.com/webhippie/go-inwx
make retool sync clean generate test
```

## Examples

### Account

[embedmd]:# (examples/account/main.go go)
```go
package main

import (
	"fmt"
	"os"

	"github.com/webhippie/go-inwx/inwx"
)

func main() {
	client, err := inwx.NewClient(
		inwx.WithUsername("owncloud"),
		inwx.WithPassword("Vater.Unser.42"),
	)

	if err != nil {
		fmt.Println("Failed to login user:", err)
		os.Exit(1)
	}

	account, _, err := client.Account.Info()

	if err != nil {
		fmt.Println("Failed to get account:", err)
		os.Exit(1)
	}

	fmt.Printf("Username: %s\n", account.Username)
	fmt.Printf("Salutation: %s\n", account.Salutation)
	fmt.Printf("Firstname: %s\n", account.Firstname)
	fmt.Printf("Lastname: %s\n", account.Lastname)
	fmt.Printf("Company: %s\n", account.Company)
}
```

## Security

If you find a security issue please contact thomas@webhippie.de first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

* [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```
