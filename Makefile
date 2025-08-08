
.PHONY: build docker kind deploy

build:
	@echo ">> Building binary"
	go build -o hello ./cmd/server

docker:
	@echo ">> Building docker image"
	docker build -t hello-website:latest .

kind: docker
	@echo ">> Loading image into kind"
	kind load docker-image hello-website:latest

deploy:
	@echo ">> Applying Kubernetes manifests"
	kubectl apply -f k8s/
