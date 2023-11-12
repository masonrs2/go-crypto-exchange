build-dev:
	go build -o bin/exchange

run-dev: build-dev
	./bin/exchange

test: 
	go test -v ./...