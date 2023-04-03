FROM golang:1.18.3-alpine3.16

WORKDIR /app

COPY client/go.mod .
COPY client/go.sum .
RUN go mod download

COPY ./client .

RUN go build -o /goapp ./cmd

EXPOSE 8089

CMD ["/goapp"]