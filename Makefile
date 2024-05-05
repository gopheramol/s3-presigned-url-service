run:
    go run main.go

build:
    go build -o s3-presigned-url-service main.go

test:
    go test ./...
