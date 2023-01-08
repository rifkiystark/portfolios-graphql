## build stage
FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download -x

COPY . .

WORKDIR /app/cmd

RUN go build -o main .

## final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/cmd/main .

EXPOSE 8080

CMD ["./main"]