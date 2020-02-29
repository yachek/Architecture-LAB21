default: out/example

rebuild: clean out/example

clean:
	rmdir /q /s out
test:
	go vet && go test

out/example: implementation.go cmd/example/main.go
	mkdir out
	FOR /F "tokens=* USEBACKQ" %%F IN (`git describe`) DO SET var=%%F
	echo package main> cmd/example/version.go
	echo const (version = ^"%var%^")>> cmd/example/version.go
	go build -o out/example ./cmd/example