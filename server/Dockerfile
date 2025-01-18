FROM golang:1.23.4-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY wait-for.sh .

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/wait-for.sh", "postgres:5432", "--" ]