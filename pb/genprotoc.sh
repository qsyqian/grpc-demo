protoc --go_out=./ --go_opt=paths=source_relative \
       --proto_path=./ \
       --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
       --grpc-gateway_out=./ --grpc-gateway_opt=paths=source_relative \
       ./hello/hello.proto