.PHONY: all clean generate details html install test update

all: clean install

clean:
	git clean -fX

generate:
	go generate ./...

details: coverage.test
	go tool cover -func=$<

html: coverage.test
	go tool cover -html=$<

install: test
	go install ./cmd/optgen

test: generate
	go test -cover -count=1 ./...

update:
	go get -u ./...

%.txt %.test: generate
	go test -cover -count=1 -coverprofile=$@ ./...
