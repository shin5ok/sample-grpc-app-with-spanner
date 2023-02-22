
.PHONY: go
go:
	protoc -Iproto --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false proto/simple.proto

.PHONY: doc
doc:
	protoc -I. -Iproto --doc_out=. --doc_opt=markdown,grpc-simple.md proto/*.proto