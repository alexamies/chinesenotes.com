FROM golang:1.9-stretch

ADD /go/src/cnweb /go/src/cnweb
RUN go get github.com/go-sql-driver/mysql
RUN go build cnweb
RUN go install cnweb

FROM debian:stretch
COPY --from=0 /go/bin/cnweb .
RUN mkdir -p /log
CMD ["./cnweb"]