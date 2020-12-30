default: build_server build_client

build_server:
	go build -o bin/k8s-client-api cmd/main.go

build_client:
	cd web && npm run build
