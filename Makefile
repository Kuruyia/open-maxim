.PHONY: build test clean docker-build docker-run linter lint format formatter lint-fix fmt fmt-check

APP_NAME = maxim_server
MAIN_PATH = ./cmd/server
GOLANGCI_LINT_VERSION = v2.14.0

build:
	go build -o bin/$(APP_NAME) $(MAIN_PATH)

test:
	go test -race -shuffle=on -count=1 ./...

clean:
	rm -rf bin/

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run --rm -ti -p8080:80 -p2002:2002 \
		-e MAXIM_CHAT_IP=localhost -e MAXIM_ADVERT_ADDR=http://localhost:8080 \
		$(APP_NAME):latest

./bin/golangci-lint:
	curl -sSfL https://golangci-lint.run/install.sh | sh -s $(GOLANGCI_LINT_VERSION)

linter: ./bin/golangci-lint
formatter: ./bin/golangci-lint

lint: linter
	./bin/golangci-lint config verify
	./bin/golangci-lint run ./...

lint-fix: linter
	./bin/golangci-lint run ./... --fix

format: formatter
	./bin/golangci-lint fmt ./...

fmt: format

fmt-check: formatter
	./bin/golangci-lint fmt --diff ./...

check: lint fmt-check
	go vet ./...
	go mod tidy -diff
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
