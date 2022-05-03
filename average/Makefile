build:
	protoc -I ./averagepb --go_opt=module=average --go_out=. --go-grpc_opt=module=average --go-grpc_out=. ./averagepb/*.proto
	go build -o ./bin/server ./server
	go build -o ./bin/client ./client

clean:
	rm ./averagepb/average.pb.go 
	rm ./averagepb/average_grpc.pb.go
	rm ./bin/server
	rm ./bin/client

up:
	./bin/server

avg:
	./bin/client