default: out/example

rebuild: clean out/example

clean:
	rmdir /q /s out
test:
	go vet && go test

out/example: implementation.go cmd/example/main.go
	mkdir out

	go build -o out/example ./cmd/example