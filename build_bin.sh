#!/bin/sh

gox -os="linux" -arch="arm" \
    -tags pi \
    -ldflags "-X github.com/gbbirkisson/rpi.revision=${TRAVIS_COMMIT} -X github.com/gbbirkisson/rpi.version=${TRAVIS_TAG}" \
    -output="{{.Dir}}_{{.OS}}_{{.Arch}}" \
    -verbose ./cmd/rpi-server

gox -os="linux" -arch="amd64 arm" \
    -ldflags "-X github.com/gbbirkisson/rpi.revision=${TRAVIS_COMMIT} -X github.com/gbbirkisson/rpi.version=${TRAVIS_TAG}" \
    -output="{{.Dir}}_{{.OS}}_{{.Arch}}" \
    -verbose ./cmd/rpi-client