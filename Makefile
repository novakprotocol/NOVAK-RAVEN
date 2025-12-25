build:
	go build -o bin/novakd ./cmd/novakd
run: build
	./bin/novakd
