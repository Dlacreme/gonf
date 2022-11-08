SHELL = /bin/bash

SRC_FOLDER = gonf
MAIN = $(SRC_FOLDER)/conf.go

.PHONY: demo
demo:
	(cd demo;go run demo.go)

.PHONY: test
test:
	go test -v
