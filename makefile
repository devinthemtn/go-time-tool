BINARY_NAME=mytime

.PHONY: build run test clean

build: main.go
		GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin main.go
		GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux main.go
		GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows main.go

run: 
		go run main.go
		# ./${BINARY_NAME}

test:
		go test -v ./...

clean:
	go clean
	rm bin/${BINARY_NAME}-darwin
	rm bin/${BINARY_NAME}-linux
	rm bin/${BINARY_NAME}-windows

install: 
	cp ./bin/mytime-darwin ~/bin/mdate
	chmod +x ~/bin/mdate
