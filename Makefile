build-prometheus:
	docker compose -f ./prometheus/docker-compose.yml up --build -d
build-proxy-server:
	docker compose -f ./proxy-server/docker-compose.yml up --build -d