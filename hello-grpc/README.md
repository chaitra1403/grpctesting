# GRPC-testing

REQUIRED .proto file

create _grpc.pb.go and .pb.go file using

    `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/hello.proto`

create mockserver using:

    `mockgen -source=proto/hello_grpc.pb.go -destination=mock_hello.pb.go -package=hello`


TO DO:

given .proto file, .go files

generate stub files using protoc

create mockserver.pb.go using stub files using mockgen

generate test files USE ROOST-GPT(protofile,mockserverstub) returns(test_files)
