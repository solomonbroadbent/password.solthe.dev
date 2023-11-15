protoc -I project/protos/ \
	--go_out=project/generated \
	--go-grpc_out=project/generated \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	project/protos/password-generator.proto
