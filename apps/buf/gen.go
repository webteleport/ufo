package main

//go:generate go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./greet/v1/greet.proto
//go:generate goimports -w .
