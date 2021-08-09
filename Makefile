build:
	echo "Building the binaries"
	go build -o . ./main.go

test:
	echo "Running unit tests"
	go test -short -v ./...

format:
	goimports -l -w .

tools:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/golangci/golangci-lint/cmd/golangci-lint

all:
	format test build