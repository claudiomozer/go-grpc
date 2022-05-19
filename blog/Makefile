build:
	protoc -I ./proto --go_opt=module=blog --go_out=. --go-grpc_opt=module=blog --go-grpc_out=. ./proto/*.proto

clean:
	rm ./proto/blog.pb.go 
	rm ./proto/blog_grpc.pb.go
