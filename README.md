# rpi

A gRPC server for remote IO operations on the RaspberryPi + cli tool to call it

## TODO

* GPIO
    * ~~Get available pins~~
    * ~~Set pin direction~~
    * Read pin values
    * Write pin values
    * Pulldown / Pullup / ActiveLow

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

## Using the client in go

```go
package main

import (
	"context"
	"fmt"
	"os"

	rpi "github.com/gbbirkisson/rpi/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to rpi-server: %v\n", err)
		os.Exit(1)
	}
	client := rpi.NewGPIOClient(conn)
	res, err := client.Pins(context.Background(), &rpi.Void{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid response from the server: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(res)
}
```

## Using another languages

Generate a client with `protoc` from [.proto](./proto) files.

## Libraries used

* [embd](https://github.com/kidoman/embd)
* [cobra](https://github.com/spf13/cobra)