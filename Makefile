balena:
	@git push balena master

docker:
	@docker build -t rpi-server .

docker-run:
	@docker run -it --rm -p 8000:8000 rpi-server