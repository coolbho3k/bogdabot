# Makefile

.PHONY: build
build:
	# Go binary
	mkdir -p bin
	go build -o bin/main cmd/main.go

.PHONY: run
run: build
	bin/main

.PHONY: clean
clean:
	rm -rf bin/*

.PHONY: default
default: build
