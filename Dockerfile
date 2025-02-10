FROM golang:1.23.3 AS builder

WORKDIR /app
COPY . . 

RUN CGO_ENABLED=0 go build  -o server -buildvcs=false ./cmd/ordersystem 

FROM scratch
WORKDIR /
COPY --from=builder /app/server .

EXPOSE 8000  
EXPOSE 50051 
EXPOSE 8080  

CMD ["./server"]
