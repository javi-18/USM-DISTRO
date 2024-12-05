FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o continente-folder .

EXPOSE 5003

CMD ["./continente-folder"]
