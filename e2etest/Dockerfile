# Docker file for chinesenotes.com end-to-end test
FROM golang:1.13 as builder

WORKDIR /app

COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server
FROM alpine:3
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/server /server
COPY --from=builder /app/static.tar.gz /static.tar.gz
RUN tar -xzf static.tar.gz
CMD ["/server"]