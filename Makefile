.PHONY: run
run:
	go run ./cmd/sketch/main.go

.PHONY: build
build:
	go build -o sketch ./cmd/sketch/main.go

.PHONY: lint
lint:
	golangci-lint run \
		--config=.golangci.pipeline.yaml \
		--sort-results \
		--max-issues-per-linter=1000 \
		--max-same-issues=1000 \
		./...
