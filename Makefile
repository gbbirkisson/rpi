.PHONY: proto

VVAR:=github.com/gbbirkisson/rpi.Version

install-cli:
	@cd ./cmd/rpi-client && go install -ldflags "-X ${VVAR}=$(shell git rev-parse HEAD)"

install-server:
	@cd ./cmd/rpi-server && go install -ldflags "-X ${VVAR}=$(shell git rev-parse HEAD)"

proto:
	@cd ./proto && protoc --go_out=plugins=grpc,import_path=rpi:. *.proto

balena:
	@git push balena master

docker:
	@docker build -t rpi-server .

docker-run:
	@docker run -it --rm -p 8000:8000 rpi-server