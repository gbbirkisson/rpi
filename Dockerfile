FROM golang:1.9-stretch as builder

COPY . $GOPATH/src/github.com/gbbirkisson/rpi
WORKDIR $GOPATH/src/github.com/gbbirkisson/rpi/cmd/rpi-server

RUN go get && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags local -ldflags "-X github.com/gbbirkisson/rpi.Version=$(git rev-parse HEAD)" -a -installsuffix cgo -o /go/bin/rpi-server

FROM alpine

COPY --from=builder /go/bin/rpi-server /rpi-server

CMD ["/rpi-server", "--gpio"]