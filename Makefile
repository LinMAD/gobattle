.DEFAULT_GOAL := pre_build

lint:
	find . -type d -not -path "*/vendor/*" | xargs -L 1 golint

test:
	go test -cover -race `go list ./... | grep -v /vendor/ | grep -v /cmd/`

pre_build: lint test
	$(info Linting and testing)

build:
	go build .