install:
	go install -v github.com/incu6us/goimports-reviser/v3@latest
	go install github.com/bufbuild/buf/cmd/buf@latest

generate:
	buf generate
