SHELL := /bin/bash
BASEDIR = $(shell pwd)

all: gotool
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
clean:
	rm -f Go-redis
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
gotool:
	gofmt -w .
help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make ca - generate ca files"

.PHONY: clean gotool help