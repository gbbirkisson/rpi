FROM golang:1.9-stretch as builder

ENV NGROK_DL="https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip" \
    GOOS="linux" \
    GOARCH="amd64" \
    GOTAGS=""

RUN apt-get update && apt-get install -yq unzip && wget ${NGROK_DL} -O /tmp/ngrok.zip && \
    unzip /tmp/ngrok.zip -d /tmp/ngrok && \
    chmod +x /tmp/ngrok/ngrok

COPY . $GOPATH/src/github.com/gbbirkisson/rpi
WORKDIR $GOPATH/src/github.com/gbbirkisson/rpi/cmd/rpi-server

RUN go get && CGO_ENABLED=0 go build ${GOTAGS} -ldflags "-X github.com/gbbirkisson/rpi.Version=$(git rev-parse HEAD)" -a -installsuffix cgo -o /tmp/rpi-server

FROM alpine

COPY --from=builder /tmp/ngrok/ngrok /usr/bin/ngrok
COPY --from=builder /tmp/rpi-server /usr/bin/rpi-server

CMD ["rpi-server"]