FROM golang:1.16.4-alpine3.13 AS builder
ENV GOPATH=/
WORKDIR /src
COPY ./ ./
RUN go mod download
RUN env GOOS=linux go build -ldflags="-s -w" -o bin/main cmd/app/main.go


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /src/bin/main .
COPY --from=builder /src/.env .
CMD ["./main"]