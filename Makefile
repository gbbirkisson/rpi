.PHONY: proto

install-cli:
	@cd ./cmd/rpi-client && go install

install-server:
	@cd ./cmd/rpi-server && go install

proto:
	@cd ./proto && protoc --go_out=plugins=grpc,import_path=rpi:. *.proto

balena:
	@git push balena master

docker:
	@docker build -t rpi-server .