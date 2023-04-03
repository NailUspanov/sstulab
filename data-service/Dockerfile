FROM golang:1.18.3-alpine3.16

WORKDIR /app

COPY data-service/go.mod .
COPY data-service/go.sum .
RUN go mod download

COPY ./data-service .

RUN go build -o /goapp ./cmd

EXPOSE 8000

CMD ["/goapp"]