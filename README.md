# stripego

![CodeQL][01]
![Go][02]
![Linter][03]
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)][04]
[![Codecov][05]][06]
[![Go Reference][07]][08]

Made stripe integration easy based on the official
[Stripe Go](https://github.com/stripe/stripe-go/) client library.

## Requirements

Go `1.19` or later

## Installation

Use Go Modules in your project:

```
go mod init <project>
```

Then, reference stripego in your code:

```
import (
	"github.com/pilinux/stripego"
)
```

Add the missing dependencies by tidying up `go.mod` file:

```
go mod tidy
```

## Features

### PaymentIntent

- [x] create a new PaymentIntent object
- [x] update the amount of an existing PaymentIntent object
- [x] update the payment method of an existing PaymentIntent object
- [x] cancel an existing PaymentIntent object

### Transfer

- [x] transfer balance to a connected Stripe account

### Transaction

- [x] get details of a balance transaction

## Usage

Please check the test files.

[01]: https://github.com/pilinux/stripego/actions/workflows/codeql-analysis.yml/badge.svg
[02]: https://github.com/pilinux/stripego/actions/workflows/go.yml/badge.svg
[03]: https://github.com/pilinux/stripego/actions/workflows/golangci-lint.yml/badge.svg
[04]: LICENSE
[05]: https://codecov.io/gh/pilinux/stripego/branch/main/graph/badge.svg?token=83H0G5TRH7
[06]: https://codecov.io/gh/pilinux/stripego
[07]: https://pkg.go.dev/badge/github.com/pilinux/stripego
[08]: https://pkg.go.dev/github.com/pilinux/stripego
