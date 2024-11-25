# GRPC-testing

1. clone this git repository
2. delete hello.pb.go, hello_grpc.pb.go, mock_hello.pb.go
3. now in root dir of the hello-grpc folder create pb.go files
4. create mockserver file using mockgen command

once all the above is ready without any issues first run server.go by using
`go run server/server.go`

#### you must see similar output in terminal

<img width="461" alt="Screenshot 2024-11-25 at 2 23 57 PM" src="https://github.com/user-attachments/assets/a3d5e47d-51bd-4427-a19e-ed588fbfe833">

#### open different terminal while the server is still running and run

`go run client/client.go`

#### you must see similar output in terminal

<img width="477" alt="Screenshot 2024-11-25 at 2 25 56 PM" src="https://github.com/user-attachments/assets/d87f31e8-c44d-4141-9dd2-0541d2b05b8d">

#### once client says hello the server acknowledge by 

<img width="507" alt="Screenshot 2024-11-25 at 2 26 36 PM" src="https://github.com/user-attachments/assets/f63e0b1d-2710-4340-86cc-317ab02490b6">

#### once both server and client is running successfully you can run the test

`go test -v server/server_test.go`
`go test -v client/client_test.go`

*** YOU DO NOT NEED ANY ACTIVE SERVER RUNNING TO RUN THIS TEST ***



### Helper commends

create _grpc.pb.go and .pb.go file using

    `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/hello.proto`

create mockserver using:

    `mockgen -source=proto/hello_grpc.pb.go -destination=mock_hello.pb.go -package=hello`




# TO DO in roostgpt:

given .proto file, .go files

generate stub files using protoc

create mockserver.pb.go using stub files using mockgen

generate test files using ROOST-GPT

input(protofile,mockserverstub) 
returns(test_files)
