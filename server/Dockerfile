FROM golang:1.18-alpine
RUN mkdir src/app
WORKDIR /src/app
COPY . .
RUN go mod tidy
EXPOSE 8080
CMD go run cmd/app/main.go