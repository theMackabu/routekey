.ONESHELL:
.PHONY: ../testing

build:
	rm -rf ../testing/routekey
	go get .
	go build .
	mv routekey ../testing