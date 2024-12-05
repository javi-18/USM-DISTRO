FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o logistica ./logistica.go

EXPOSE 50051

CMD ["./logistica"]