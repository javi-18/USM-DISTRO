FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o servidores-hextech .

EXPOSE 5002 5003 5004

CMD ["./servidores-hextech"]
