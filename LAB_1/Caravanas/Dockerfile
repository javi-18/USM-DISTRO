FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o caravanas ./caravanas.go

EXPOSE 50052

CMD ["./caravanas"]
