FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
EXPOSE 8000
ENTRYPOINT [ "go", "run", "main.go" ]
