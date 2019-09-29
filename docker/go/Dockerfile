FROM golang:1.13-buster

ENV GO111MODULE=on
WORKDIR /cnweb
COPY /go/src/cnweb .
RUN go build
RUN apt-get update
RUN apt-get install -y ca-certificates
RUN mkdir -p data
COPY data/words.txt data/.
CMD ["./cnweb"]