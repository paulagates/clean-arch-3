FROM golang:1.23.3 AS builder

WORKDIR /app
COPY . . 
COPY .env .env

RUN CGO_ENABLED=0 go build  -o server ./cmd/ordersystem 

FROM scratch
WORKDIR /
COPY --from=builder /app/server .
COPY --from=builder /app/.env .env

EXPOSE 50051
EXPOSE 8080
EXPOSE 8000

CMD ["./server"]
