protoc --go_out=./gate ./pb/gate/*.proto
protoc --go_out=./grpc_server ./pb/grpc_server/*.proto
protoc --go-grpc_out=./grpc_server ./pb/grpc_server/*.proto
