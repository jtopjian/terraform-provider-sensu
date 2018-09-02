# terraform-provider-sensu

[![Build Status](https://travis-ci.org/jtopjian/terraform-provider-sensu.svg?branch=master)](https://travis-ci.org/jtopjian/terraform-provider-sensu)


Sensu 2.0 resource provider for Terraform

> Sensu 2.0 and the Sensu 2.0 API are still in beta stages.
>
> This provider should not be considered stable nor does it
> guarantee any type of backwards compatibility.

## Prerequisites

* [Terraform][1]

## Terraform Configuration Example

```hcl
provider "sensu" {
	api_url      = "http://127.0.0.1:8080"
	username     = "admin"
	password     = "password"
	environment  = "default"
	organization = "default"
}

resource "sensu_check" "check_1" {
	name     = "check_1"
	command  = "/bin/foo"
	interval = 600

	subscriptions = [
		"foo",
		"bar",
	]
}
```

## Installation

### Using a Pre-Built Binary

Downloading and installing a pre-compiled `terraform-provider-sensu` release
is the recommended method of installation since it requires no additional tools
or libraries to be installed on your workstation.

1. Visit the [releases][2] page and download the latest release for your target
   architecture.

2. Unzip the downloaded file and copy the `terraform-provider-sensu` binary
   to a designated directory as described in Terraform's [plugin installation
   instructions][3].

### Building from Source

> Note: Terraform requires Go 1.9 or later to successfully compile.

1. Follow these [instructions][4] to setup a Golang development environment.
2. Run:

```shell
$ go get -v -u github.com/jtopjian/terraform-provider-sensu
$ cd $GOPATH/src/github.com/jtopjian/terraform-provider-sensu
$ make build
```

You should now have a `terraform-provider-sensu` binary located at
`$GOPATH/bin/terraform-provider-sensu`. Copy this binary to a designated
directory as described in Terraform's [plugin installation instructions][3]

## Development

This project is using [`govendor`][5] for the time being. In the future,
it'll probably move to the official Go module support.

## Documentation

Full documentation can be found in the [`docs`][6] directory.

[1]: http://terraform.io
[2]: https://github.com/jtopjian/terraform-provider-sensu/releases
[3]: https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin
[4]: https://golang.org/doc/install
[5]: https://github.com/kardianos/govendor
[6]: /docs
