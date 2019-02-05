#!/bin/sh

gox -osarch="linux/arm" \
    -tags pi \
    -ldflags "-X github.com/gbbirkisson/rpi.revision=${TRAVIS_COMMIT} -X github.com/gbbirkisson/rpi.version=${TRAVIS_TAG}" \
    -output="{{.Dir}}_{{.OS}}_{{.Arch}}" \
    -verbose ./cmd/rpi-server

gox -osarch="linux/amd64 linux/arm darwin/amd64 windows/amd64" \
    -ldflags "-X github.com/gbbirkisson/rpi.revision=${TRAVIS_COMMIT} -X github.com/gbbirkisson/rpi.version=${TRAVIS_TAG}" \
    -output="{{.Dir}}_{{.OS}}_{{.Arch}}" \
    -verbose ./cmd/rpi-client