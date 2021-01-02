default: fmt test

test:
	go test .

fmt:
	go run golang.org/x/tools/cmd/goimports -w .

.PHONY: default test fmt
