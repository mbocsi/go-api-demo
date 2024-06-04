FROM golang:1.22.3-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . . 

RUN go build -o /goapi ./cmd/api

CMD ["/goapi"]
