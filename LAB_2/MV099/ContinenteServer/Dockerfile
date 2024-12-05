FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o continente-server .

EXPOSE 5005

CMD ["./continente-server"]
