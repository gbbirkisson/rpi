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

## TODO

* rpi-client
    * gpio
        * Write
        * Read
        * Freq
        * DutyCycle
        * Detect
        * EdgeDetected

## Pi Camera

RESIN_HOST_CONFIG_gpu_mem 128
RESIN_HOST_CONFIG_start_x 1

modprobe bcm2835-v4l2

## Using another languages

Generate a client with `protoc` from [.proto](./proto) files.

## Libraries used

* [go-rpio](https://github.com/stianeikeland/go-rpio)
* [cobra](https://github.com/spf13/cobra)