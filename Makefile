default: out/example

rebuild: clean out/example

clean:
	rmdir /q /s out
test:
	go vet && go test

out/example: implementation.go cmd/example/main.go
	echo package main> cmd/example/version.go
	echo const (version = %1)>> cmd/example/version.go | git describe
	go build -o out/example ./cmd/example