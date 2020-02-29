default: out/example

rebuild: clean out/example

clean:
	rmdir /q /s out
test:
	go vet && go test
define VERSION_BODY
package main

const BuildVersion = "$(shell git describe)"
endef
export VERSION_BODY
out/example: implementation.go cmd/example/main.go
	mkdir out
	echo "$$VERSION_BODY" > version.go
	go build -o out/example ./cmd/example