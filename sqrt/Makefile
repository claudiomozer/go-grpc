build:
	protoc -I ./sqrtpb --go_opt=module=sqrt --go_out=. --go-grpc_opt=module=sqrt --go-grpc_out=. ./sqrtpb/*.proto
	go build -o ./bin/server ./server
	go build -o ./bin/client ./client

clean:
	rm ./sqrtpb/sqrt.pb.go 
	rm ./sqrtpb/sqrt_grpc.pb.go
	rm ./bin/server
	rm ./bin/client

up:
	./bin/server

sqrt:
	./bin/client