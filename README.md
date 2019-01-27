# rpi

<p align="center">
<a href="https://github.com/gbbirkisson/rpi/releases/latest" target="_blank"><img src="https://img.shields.io/github/release/gbbirkisson/rpi.svg?style=flat-square"/></a>
<a href="https://travis-ci.org/gbbirkisson/rpi" target="_blank"><img src="https://img.shields.io/travis/gbbirkisson/rpi/master.svg?style=flat-square"/></a>
<a href="https://goreportcard.com/report/github.com/gbbirkisson/rpi" target="_blank"><img src="https://goreportcard.com/badge/github.com/gbbirkisson/rpi?style=flat-square"/></a>
<a href="https://godoc.org/github.com/gbbirkisson/rpi" target="_blank"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"/></a>
<a href="./LICENSE" target="_blank"><img src="https://img.shields.io/badge/license-Apache%202.0-blue.svg?style=flat-square"/></a>
</p>

This package helps you to develop software for the Raspberry Pi that does IO operations on the Raspberry PI. It enables you to develop your code locally on any type of architecture, by using gRPC to control a Raspberry PI remotely. 

This is very convenient in conjunction with services like balena.io. This makes developing applications for the RaspberryPi extremely easy. Once your software is ready, you have the option of continuing to use gRPC calls, or switch over to a local version of the interfaces to compile a binary that runs directly on the RaspberryPi.

If you have any suggestion or comments, please feel free to open an issue on this GitHub page.

## Installing

#### Pushing server to balena.io

Add your balena remote to the git repo:

```bash
git add remote balena <URL>
```

```bash
make balena
```

## Cli usage

Run the program using help to get documentation

```
rpi-client help
```

## Server modules

### GPIO

To enable on balena.io:

* Device Service variables
    * `RPI_GPIO`: `true`

### Pi Camera

To enable on balena.io:

* Device Configuration
    * `RESIN_HOST_CONFIG_gpu_mem`: `128`
    * `RESIN_HOST_CONFIG_start_x`: `1`
* Device Service variables
    * `RPI_PICAM`: `true`
    * `RPI_MODPROBE`: `bcm2835-v4l2`

### Ngrok tunnel

* Device Service variables
    * `RPI_NGROK`: `true`
    * `RPI_NGROK_TOKEN`: `<your ngrok token>`

## Using another languages

Generate a client with `protoc` from [./pkg/proto/*.proto](./pkg/proto) files.
