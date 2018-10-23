FROM golang:1.11

ENV DIR=/code
RUN mkdir $DIR

WORKDIR $DIR
ADD . .
RUN go mod download
RUN go build .

ENTRYPOINT ./http
