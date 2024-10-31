FROM golang:1.23.1 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Compilar el binario principal
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app

FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/main ./
# COPY --from=build /templates /templates

EXPOSE 8000

CMD ["./main"]
