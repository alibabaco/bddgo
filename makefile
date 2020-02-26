
.PHONY: build
build:
	go build ./...


.PHONY: install
install:
	go install ./...


.PHONY: test
test: build
	go test
