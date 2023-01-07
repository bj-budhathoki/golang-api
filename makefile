BINARY_NAME=main.out

all: build test

build:
	go build -o ${BINARY_NAME} main.go

test:
	go test -v main.go

run:
	go build -o ${BINARY_NAME} main.go
	 ./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

hello:
	echo "Hello"
compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go
