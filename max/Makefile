build:
	protoc -I ./maxpb --go_opt=module=max --go_out=. --go-grpc_opt=module=max --go-grpc_out=. ./maxpb/*.proto
	go build -o ./bin/server ./server
	go build -o ./bin/client ./client

clean:
	rm ./maxpb/max.pb.go 
	rm ./maxpb/max_grpc.pb.go
	rm ./bin/server
	rm ./bin/client

up:
	./bin/server

max:
	./bin/client