# run :
# 	docker-compose -f docker-compose.dev.yml up

# down :
# 	docker-compose -f docker-compose.dev.yml down
build:     
	go build -o ./go-api ./server.go      
run:     
	go run server.go 
runcontainer:     
	docker run -d -p 8000:8000 --env-file ./.env go-api 
compile:     
	echo "Compiling for every OS and Platform"     
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go     
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go     
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go 