build:
	protoc -I ./primepb --go_opt=module=prime --go_out=. --go-grpc_opt=module=prime --go-grpc_out=. ./primepb/*.proto
	go build -o ./bin/server ./server
	go build -o ./bin/client ./client

clean:
	rm ./primepb/prime.pb.go 
	rm ./primepb/prime_grpc.pb.go
	rm ./bin/server
	rm ./bin/client

up:
	./bin/server

decompose:
	./bin/client