# rpi

<p align="center">
<a href="https://github.com/gbbirkisson/rpi/releases/latest" target="_blank"><img src="https://img.shields.io/github/release/gbbirkisson/rpi.svg?style=flat-square"/></a>
<a href="https://travis-ci.org/gbbirkisson/rpi" target="_blank"><img src="https://img.shields.io/travis/gbbirkisson/rpi/master.svg?style=flat-square"/></a>
<a href="https://goreportcard.com/report/github.com/gbbirkisson/rpi" target="_blank"><img src="https://goreportcard.com/badge/github.com/gbbirkisson/rpi?style=flat-square"/></a>
<a href="https://godoc.org/github.com/gbbirkisson/rpi" target="_blank"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"/></a>
<a href="./LICENSE" target="_blank"><img src="https://img.shields.io/badge/license-Apache%202.0-blue.svg?style=flat-square"/></a>
</p>

This package helps you to develop software for the RaspberryPi that does IO operations on the RaspberryPi. It enables you to develop your code locally on any type of architecture, by using gRPC to control a Raspberry PI remotely. This is very convenient in conjunction with services like [balena.io](https://www.balena.io/). 

This makes developing applications for the RaspberryPi extremely easy. Once your software is ready, you have the option of continuing to use gRPC calls, or switch over to a local version of the interfaces to compile a binary that runs directly on the RaspberryPi.

If you have any suggestion or comments, please feel free to open an issue on this GitHub page.

## Installing

### Client / Server

There are 2 ways of installing the binaries:

* Download the binaries [latest release](https://github.com/gbbirkisson/rpi/releases/latest)
* Install from source: `go get github.com/gbbirkisson/rpi`

### Pushing server to balena.io

Take a look at the [rpi-balena](https://github.com/gbbirkisson/rpi-balena) project to see how to use balena.io

## Configuration

Both server and client can be configured with:
1. Flags
2. Environment
3. Configuration file

The presidence is in that order, i.e flags override environment that overrides the configuration file.

### Flags

Use `-h` flag to see available flags for `rpi-client` and `rpi-server`.

### Configuration files

Configuration files can be in the following formats:
* json
* toml
* yaml
* hcl

Locations of those files are:
* Server: `/etc/rpi-server/config.[format]`
* Client: `~/.rpi-client.[format]`

To generate the default configuration files (yaml in this example) do:

```bash
# For server
$ touch /etc/rpi-server/config.yaml
$ rpi-server config write

# For client
$ touch ~/.rpi-client.yaml
$ rpi-client config write
```

The resulting client configuration file (`~/.rpi-client.yaml`) would be something like:

```yaml
server:
  host: 127.0.0.1
  port: 8000
  timeout: 5000
picam:
  viewer:
  - feh
  - -x
  - '-'
```

### Environmental variables

Environmental variables mirror the configuration files. All variables have the prefix `RPI_`. So for example if you want to set the client timeout you can set it with the environment variable `RPI_SERVER_TIMEOUT=3000`

## Using another languages

Generate a client for your language of choice with `protoc` using [./pkg/proto/*.proto](./pkg/proto) files.
