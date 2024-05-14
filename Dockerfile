## Build stage
FROM golang:1.21.10-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
# RUN go build -o main main.go

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main . 
COPY pkg/db/migration ./migration
COPY app.env /app/


EXPOSE 8080
CMD ["/app/main"]
# ENTRYPOINT [ "/app/start.sh" ]
