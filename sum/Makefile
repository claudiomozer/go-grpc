build:
	protoc -I ./sumpb --go_opt=module=sum --go_out=. --go-grpc_opt=module=sum --go-grpc_out=. ./sumpb/*.proto
	go build -o ./bin/server ./server
	go build -o ./bin/client ./client

clean:
	rm ./sumpb/sum.pb.go 
	rm ./sumpb/sum_grpc.pb.go
	rm ./bin/server
	rm ./bin/client

up:
	./bin/server

sum:
	./bin/client