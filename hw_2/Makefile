LINE ?= ""

.PHONY: run
run: vendor vet golint test
	@go run main.go -line=$(LINE)

.PHONY: test
test:
	@go test -v -i $(shell go list ./... | grep -v '/vendor/')
	@go test -v $(shell go list ./... | grep -v '/vendor/')

.PHONY: vendor
vendor:
	@go mod vendor

.PHONY: vet
vet:
	@go vet $(shell go list ./... | grep -v '/vendor/')

.PHONY: golint
golint:
	@go get -u golang.org/x/lint/golint
	@golint $(shell go list ./... | grep -v '/vendor/')