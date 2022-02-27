gen-common-pb:
	rm -rf ./pb/common
	mkdir -p ./pb/common
	protoc --proto_path=./proto/common \
	--go_out=./pb/common \
	--go_opt=paths=source_relative \
 	--go-grpc_out=./pb/common \
 	--go-grpc_opt=paths=source_relative \
 	--go-grpc_opt=require_unimplemented_servers=false \
 	proto/common/*.proto

gen-wallet-service-pb:
	rm -rf ./pb/wallet-service
	mkdir -p ./pb/wallet-service
	protoc --proto_path=./proto/wallet-service \
	--go_out=./pb/wallet-service \
	--go_opt=paths=source_relative \
 	--go-grpc_out=./pb/wallet-service \
 	--go-grpc_opt=paths=source_relative \
 	--go-grpc_opt=require_unimplemented_servers=false \
 	proto/wallet-service/*.proto

gen-all: gen-wallet-service-pb gen-common-pb