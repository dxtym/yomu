FROM golang:1.23.4-alpine3.20 AS build
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

FROM alpine:3.20
WORKDIR /app
COPY app.env .
COPY --from=build /app/main .

EXPOSE 8080
CMD [ "/app/main" ]
