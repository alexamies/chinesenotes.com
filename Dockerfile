# Docker file for chinesenotes.com web app
FROM golang:1.15.8 as builder
ADD https://api.github.com/repos/alexamies/chinesenotes-go/git/refs/heads/master version.json
RUN git clone https://github.com/alexamies/chinesenotes-go.git --branch v0.0.63
WORKDIR /go/chinesenotes-go
COPY config.yaml .
COPY data/*.txt data/
RUN go build
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o cnweb
FROM alpine:3
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/chinesenotes-go/cnweb /cnweb
COPY --from=builder /go/chinesenotes-go/config.yaml /config.yaml
COPY --from=builder /go/chinesenotes-go/data/*.txt /data/
CMD ["./cnweb"]