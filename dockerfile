FROM golang:latest
LABEL maintainer="Bob"
WORKDIR $GOPATH/src/checkout-backend
COPY . $GOPATH/src/checkout-backend
RUN go build main.go
EXPOSE 8888
ENTRYPOINT ["./main"]
