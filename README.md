[![Build](https://img.shields.io/travis/gbbirkisson/rpi/master.svg?style=flat-square)](https://travis-ci.org/gbbirkisson/rpi)
[![Go Report Card](https://goreportcard.com/badge/github.com/gbbirkisson/rpi?style=flat-square)](https://goreportcard.com/report/github.com/gbbirkisson/rpi)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/gbbirkisson/rpi)
[![Release](https://img.shields.io/github/release/gbbirkisson/rpi.svg?style=flat-square)](https://github.com/gbbirkisson/rpi/releases/latest)

# rpi

A gRPC server for remote IO operations on the RaspberryPi + cli tool to call it

## Installing

#### The client

```bash
make install-cli
```

#### The server

```bash
make install-server
```

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
    * `RPI_CAMERA`: `true`
    * `RPI_MODPROBE`: `bcm2835-v4l2`

### Ngrok tunnel

* Device Service variables
    * `RPI_NGROK`: `true`
    * `RPI_NGROK_TOKEN`: `<your ngrok token>`

## Using another languages

Generate a client with `protoc` from [.proto](./proto) files.