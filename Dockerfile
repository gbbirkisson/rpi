FROM golang:1.9-stretch as builder

ENV NGROK_DL="https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip" \
    GOOS="linux" \
    GOARCH="amd64" \
    GOTAGS=""

RUN apt-get update && apt-get install -yq unzip wget && wget ${NGROK_DL} -O /tmp/ngrok.zip && \
    unzip /tmp/ngrok.zip -d /tmp/ngrok && \
    chmod +x /tmp/ngrok/ngrok

COPY . $GOPATH/src/github.com/gbbirkisson/rpi
WORKDIR $GOPATH/src/github.com/gbbirkisson/rpi

RUN go get -t -v ./...
RUN cd cmd/rpi-client && CGO_ENABLED=0 go build -a -installsuffix cgo -o /tmp/rpi-client
RUN cd cmd/rpi-server && CGO_ENABLED=0 go build ${GOTAGS} -a -installsuffix cgo -o /tmp/rpi-server

FROM alpine

COPY --from=builder /tmp/ngrok/ngrok /usr/bin/ngrok
COPY --from=builder /tmp/rpi-client /usr/bin/rpi-client
COPY --from=builder /tmp/rpi-server /usr/bin/rpi-server

CMD ["rpi-server"]