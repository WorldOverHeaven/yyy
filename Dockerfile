FROM golang:1.20-alpine AS builder

WORKDIR /build

ADD go.mod .

ADD go.sum .

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/main /build/main


CMD ["./main"]