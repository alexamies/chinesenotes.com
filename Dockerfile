# Docker file for chinesenotes.com web app
FROM golang:1.22.1 as builder
ADD https://api.github.com/repos/alexamies/chinesenotes-go/git/refs/heads/master version.json
RUN git clone https://github.com/alexamies/chinesenotes-go.git --branch v0.0.180
WORKDIR /go/chinesenotes-go
COPY *.yaml ./
COPY data/*.txt data/
COPY data/*.tsv data/
COPY data/corpus/collections.csv ./data/corpus/
COPY index/documents.tsv index/
RUN go build
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o cnweb
FROM alpine:3
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/chinesenotes-go/cnweb /cnweb
COPY --from=builder /go/chinesenotes-go/*.yaml /
COPY --from=builder /go/chinesenotes-go/data/*.txt /data/
COPY --from=builder /go/chinesenotes-go/data/*.tsv /data/
COPY --from=builder /go/chinesenotes-go/data/corpus/collections.csv /data/corpus/
COPY --from=builder /go/chinesenotes-go/index/documents.tsv /index/
CMD ["./cnweb"]