all:
	make clean
	make build
	make run

build:
	go build -o bin/main main.go

run:
	go run main.go

clean:
	rm -rf bin/*