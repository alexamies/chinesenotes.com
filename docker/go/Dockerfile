FROM golang:1.9-stretch

ADD /go/src/cnweb /go/src/cnweb
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/sendgrid/sendgrid-go
RUN go build cnweb
RUN go install cnweb

FROM debian:stretch
RUN apt-get update
RUN apt-get install -y ca-certificates
COPY --from=0 /go/bin/cnweb .
COPY --from=0 /go/src/cnweb/webconfig.yaml .
RUN mkdir -p templates
COPY --from=0 /go/src/cnweb/templates templates/.
CMD ["./cnweb"]