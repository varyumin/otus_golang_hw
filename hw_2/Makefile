run: vendor vet test
	@go run main.go

test:
	@go test -v -i $(shell go list ./... | grep -v '/vendor/')
	@go test -v $(shell go list ./... | grep -v '/vendor/')

vendor:
	@go mod vendor

vet:
	@go vet $(shell go list ./... | grep -v '/vendor/')