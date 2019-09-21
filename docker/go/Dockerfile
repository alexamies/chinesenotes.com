FROM golang:1.12-buster

ENV GO111MODULE=on
WORKDIR /cnweb
COPY /go/src/cnweb .
RUN go build

FROM debian:buster
RUN apt-get update
COPY --from=0 /cnweb/cnweb .
COPY --from=0 /cnweb/webconfig.yaml .
RUN mkdir -p templates
COPY --from=0 /cnweb/templates templates/.
RUN mkdir -p data
COPY data/words.txt data/.
CMD ["./cnweb"]