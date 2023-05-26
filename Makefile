hello:	
	echo "Hello"

build: 
	go build -o bin/main main.go
	
run: main.go
	nodemon --exec "go run" main.go

