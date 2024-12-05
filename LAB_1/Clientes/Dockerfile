FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o clientes ./clientes.go

EXPOSE 50051

CMD ["./clientes"]