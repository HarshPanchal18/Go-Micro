FROM golang:1.23.2

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o user-service

EXPOSE 8081

CMD ["./user-service"]

# docker build -t user-service .