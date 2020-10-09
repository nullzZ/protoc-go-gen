go install protoc-gen-go-nrpc.go
protoc --go-nrpc_out=plugins=nrpc:. hello.proto