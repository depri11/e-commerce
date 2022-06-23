FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o bin/main .

CMD ["./bin/main"]