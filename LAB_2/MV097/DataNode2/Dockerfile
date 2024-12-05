FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o datanode2 .

EXPOSE 5006

CMD ["./datanode2"]
