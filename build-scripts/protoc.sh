protoc -I project/protos/ \
	--go_out=project/generated \
	--go-grpc_out=project/generated \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	project/protos/password-generator.proto

protoc -I project/protos/ \
	--grpc-web_out=frontend/generated \
	--grpc-web_opt=mode=grpcwebtext \
	--grpc-web_opt=import_style=closure \
	project/protos/password-generator.proto
