FROM golang:1.23.2

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o order-service

EXPOSE 8082

CMD ["./order-service"]

# docker build -t order-service .