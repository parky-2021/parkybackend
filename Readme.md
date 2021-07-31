Requirements
Download latest version of Go


Install protoc
https://github.com/protocolbuffers/protobuf/releases/download/v3.17.3/protoc-3.17.3-win64.zip
Set path

to generate proto files
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  proto/<filename>.proto


To run Server
go run server/main.go

To run Client
go run client/main.go
