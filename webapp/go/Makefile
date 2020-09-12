all: isuumo

isuumo: *.go
	go build -o isuumo

lint:
	docker run -v $(CURDIR):/app -v $(GOPATH)/pkg/mod:/go/pkg/mod -w /app --rm golangci/golangci-lint:latest-alpine golangci-lint run ./...
