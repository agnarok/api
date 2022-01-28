.PHONY: setup
setup:
	go mod download
	go mod tidy
	go mod vendor

