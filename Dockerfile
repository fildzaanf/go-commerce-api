FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env ./  

RUN go build -o /go-commerce-api ./cmd/main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /go-commerce-api /go-commerce-api
COPY --from=builder /app/.env .env  

EXPOSE 8080

ENTRYPOINT ["/go-commerce-api"]