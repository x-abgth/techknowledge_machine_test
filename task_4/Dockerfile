# Building stage
FROM golang:1.20.0-alpine3.17 AS builder

# maintainer info
LABEL maintainer = "Abhijith A <abhijithworkemail@gmail.com>"

WORKDIR /app

COPY . .

RUN go mod tidy

RUN cd cmd/api/ && go build -o main

# Final running stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/cmd/api/main .

COPY --from=builder /app/.env .

RUN chmod +x ./main

CMD ["./main"]

EXPOSE 3000

