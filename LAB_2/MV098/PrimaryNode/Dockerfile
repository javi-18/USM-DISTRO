FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o primary-node .

EXPOSE 5002

CMD ["./primary-node"]
