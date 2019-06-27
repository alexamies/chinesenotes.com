FROM golang:1.11-stretch

ADD /go/src/cnweb /go/src/cnweb
ADD /data/words.txt /data/words.txt
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u cloud.google.com/go/storage
RUN go get -u github.com/sendgrid/sendgrid-go
RUN go build cnweb
RUN go install cnweb

FROM debian:stretch
RUN apt-get update
RUN apt-get install -y ca-certificates
COPY --from=0 /go/bin/cnweb .
COPY --from=0 /go/src/cnweb/webconfig.yaml .
RUN mkdir -p templates
COPY --from=0 /go/src/cnweb/templates templates/.
RUN mkdir -p data
COPY --from=0 /data/words.txt data/.
CMD ["./cnweb"]