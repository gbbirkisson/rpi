# rpi

<p align="center">
<a href="https://github.com/gbbirkisson/rpi/releases/latest" target="_blank"><img src="https://img.shields.io/github/release/gbbirkisson/rpi.svg?style=flat-square"/></a>
<a href="https://travis-ci.org/gbbirkisson/rpi" target="_blank"><img src="https://img.shields.io/travis/gbbirkisson/rpi/master.svg?style=flat-square"/></a>
<a href="https://goreportcard.com/report/github.com/gbbirkisson/rpi" target="_blank"><img src="https://goreportcard.com/badge/github.com/gbbirkisson/rpi?style=flat-square"/></a>
<a href="https://godoc.org/github.com/gbbirkisson/rpi" target="_blank"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"/></a>
<a href="./LICENSE" target="_blank"><img src="https://img.shields.io/badge/license-Apache%202.0-blue.svg?style=flat-square"/></a>
</p>

This package helps you to develop software for the Raspberry Pi that does IO operations on the Raspberry PI. It enables you to develop your code locally on any type of architecture, by using gRPC to control a Raspberry PI remotely. This is very convenient in conjunction with services like balena.io. 

This makes developing applications for the RaspberryPi extremely easy. Once your software is ready, you have the option of continuing to use gRPC calls, or switch over to a local version of the interfaces to compile a binary that runs directly on the RaspberryPi.

If you have any suggestion or comments, please feel free to open an issue on this GitHub page.

## Installing

### Client / Server

There are 2 ways of installing the binaries:

* Download the binaries [latest release](https://github.com/gbbirkisson/rpi/releases/latest)
* Install from source: `go get github.com/gbbirkisson/rpi`

### Pushing server to balena.io

Take a look at the [rpi-balena](https://github.com/gbbirkisson/rpi-balena) project to see how to use balena.io

## Cli usage

Run the program using help to get documentation

```
rpi-client help
```

## Using another languages

Generate a client for your language of choice with `protoc` using [./pkg/proto/*.proto](./pkg/proto) files.
