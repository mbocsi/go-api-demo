FROM golang:1.22.3-alpine

ENV GO_ENV=PROD

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . . 

RUN go build -o /goapi ./cmd/api

CMD ["/goapi"]
