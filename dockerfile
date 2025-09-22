
FROM golang:1.25-alpine AS builder


RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o main .


FROM alpine:3.20

RUN apk add --no-cache libc6-compat

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
