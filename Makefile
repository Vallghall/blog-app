run server:
	cd server && \
	go mod tidy && \
	go run -v ./cmd/app/main.go

build-server: # Rule for building Docker image
	cd server && \
	docker build . -t blog-server

docker-run-server:
	docker run -d --name blog-server -p 8080:8080 --rm blog-server