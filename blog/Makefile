build:
	protoc -I ./proto --go_opt=module=blog --go_out=. --go-grpc_opt=module=blog --go-grpc_out=. ./proto/*.proto
	go build -o ./bin/server ./server
	go build -o ./bin/client ./client

clean:
	rm ./proto/blog.pb.go 
	rm ./proto/blog_grpc.pb.go
	rm ./bin/server
	rm ./bin/client

up:
	./bin/server

blog:
	./bin/client