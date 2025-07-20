FROM golang:1.24.3 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o frontend .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/frontend .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/pages ./pages
COPY --from=builder /app/assets ./assets

EXPOSE 8080

ENTRYPOINT ["./frontend"]