FROM golang:1.24-alpine AS builder

WORKDIR /build
COPY . .

RUN go mod tidy && go build -o ./api

EXPOSE 8080

CMD ["./api"]