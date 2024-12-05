FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o isla-file .

EXPOSE 5001

CMD ["./isla-file"]