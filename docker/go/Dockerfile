FROM golang:1.12-stretch

ENV GO111MODULE=on
WORKDIR /cnweb
COPY /go/src/cnweb .
ADD /data/words.txt /data/words.txt
RUN go build

FROM debian:stretch
RUN apt-get update
RUN apt-get install -y ca-certificates
COPY --from=0 /cnweb/cnweb .
COPY --from=0 /cnweb/webconfig.yaml .
RUN mkdir -p templates
COPY --from=0 /cnweb/templates templates/.
RUN mkdir -p data
COPY --from=0 /data/words.txt data/.
CMD ["./cnweb"]