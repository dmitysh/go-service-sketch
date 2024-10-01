.PHONY: run
run:
	go run ./cmd/sketch/main.go

.PHONY: build
build:
	go build -o sketch ./cmd/sketch/main.go
