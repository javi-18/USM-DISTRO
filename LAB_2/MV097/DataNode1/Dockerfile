FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o datanode1 .

EXPOSE 5002

CMD ["./datanode1"]
