BINARY_NAME=main.out

all: deps build test

deps:
	go get -u github.com/gin-gonic/gin@v1.7
	go get -u github.com/stretchr/testify/assert
	go get -u github.com/google/uuid


build:
	go build -o ${BINARY_NAME} server/main.go

test:
	go test -v server/main.go

run:
	go build -o ${BINARY_NAME} server/main.go
	./${BINARY_NAME}
 
clean:
	go clean
	rm ./${BINARY_NAME}