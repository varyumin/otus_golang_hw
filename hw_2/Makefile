run: vendor vet golint test
	@go run main.go

test:
	@go test -v -i $(shell go list ./... | grep -v '/vendor/')
	@go test -v $(shell go list ./... | grep -v '/vendor/')

vendor:
	@go mod vendor

vet:
	@go vet $(shell go list ./... | grep -v '/vendor/')

golint:
	@golint $(shell go list ./... | grep -v '/vendor/')