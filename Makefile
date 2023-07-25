GRPC_PROTO_PATH = ./infrastructure/grpc/props/fetch-server.proto
GRPC_GEN_PATH = .

build-proxy-server:
	docker compose -f ./proxy-server/docker-compose.yml up --build -d

grpc-generate:
	protoc \
		--go_out=${GRPC_GEN_PATH}  --go_opt=paths=source_relative \
		--go-grpc_out=$(GRPC_GEN_PATH) --go-grpc_opt=paths=source_relative \
    ${GRPC_PROTO_PATH}	