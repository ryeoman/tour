.PHONY : build run

build:
	go mod tidy
	go mod vendor
	go build -o cmd/bin/tour cmd/app/main.go

run: build
	./cmd/bin/tour