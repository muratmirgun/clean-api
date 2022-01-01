### Stage 1
FROM golang:1.17.5-alpine AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/api

### Stage 2
FROM scratch
COPY --from=builder /app/api /bin/api
ENTRYPOINT ["/bin/api"]