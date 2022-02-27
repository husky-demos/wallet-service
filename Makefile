gen-wallet-service-pb:
	mkdir -p ./pb/wallet-service
	protoc --proto_path=./proto/wallet-service \
	--go_out=./pb/wallet-service \
	--go_opt=paths=source_relative \
 	--go-grpc_out=./pb/wallet-service \
 	--go-grpc_opt=paths=source_relative \
 	proto/wallet-service/*.proto