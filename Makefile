.PHONY: check test

check:
  - "gofmt -e ./internal/.."

test:
  - "go test ./.. -v"
