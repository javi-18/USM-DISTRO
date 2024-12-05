FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir -p /app/logs

RUN go build -o jayce .

EXPOSE 5005

CMD ["./jayce"]