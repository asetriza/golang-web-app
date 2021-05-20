FROM golang:1.16.4-alpine3.13

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN env GOOS=linux go build -ldflags="-s -w" -o bin/main cmd/app/main.go

CMD ["bin/./main"]
