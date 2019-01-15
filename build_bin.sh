#!/bin/sh

gox -os="linux" -arch="arm" \
    -tags pi \
    -ldflags "-X github.com/gbbirkisson/rpi.Revision=${TRAVIS_COMMIT} -X github.com/gbbirkisson/rpi.Version=${TRAVIS_TAG}" \
    -output="{{.Dir}}_{{.OS}}_{{.Arch}}_local" \
    -verbose ./...

gox -os="linux" -arch="amd64 arm" \
    -ldflags "-X github.com/gbbirkisson/rpi.Revision=${TRAVIS_COMMIT} -X github.com/gbbirkisson/rpi.Version=${TRAVIS_TAG}" \
    -output="{{.Dir}}_{{.OS}}_{{.Arch}}" \
    -verbose ./cmd/rpi-client