FROM golang:latest
LABEL maintainer="Bob"
WORKDIR $GOPATH/src/checkout-backend
COPY . $GOPATH/src/checkout-backend
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go
EXPOSE 8888
ENTRYPOINT ["./main"]
