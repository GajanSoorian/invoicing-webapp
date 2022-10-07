# This make file provides helpful wrappers for Go build commands.

# Name of the executable.
BINARY_NAME=parallax_invoicing.out

# Instal dependencies, run unit test cases, build and run the generated binary.
all: deps test build run 

# Fetch and install all backend dependencies.
deps:
	go get github.com/gin-gonic/gin@v1.7.7
	go get github.com/google/uuid@v1.3.0
	go get github.com/joho/godotenv@v1.4.0
	go get github.com/lib/pq@v1.10.7
	go get github.com/stretchr/testify@v1.8.0

# Build and create an executable binary.
build:
	go build -o ${BINARY_NAME} invoice-service/main.go

# Run unit test cases.
test:
	go test ./... -v

# Build and run the executable binary.
run:
	go build -o ${BINARY_NAME} invoice-service/main.go
	./${BINARY_NAME}

# Delete object files and binary file.
clean:
	go clean
	rm ./${BINARY_NAME}