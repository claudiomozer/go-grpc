build:
	protoc -I ./greetpb --go_opt=module=greet --go_out=. --go-grpc_opt=module=greet --go-grpc_out=. ./greetpb/*.proto
	go build -o ./bin/server ./server
	go build -o ./bin/client ./client

clean:
	rm ./greetpb/greet.pb.go 
	rm ./greetpb/greet_grpc.pb.go
	rm ./bin/server
	rm ./bin/client