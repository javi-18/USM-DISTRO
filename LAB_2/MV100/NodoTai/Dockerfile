FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o nodo_tai .

EXPOSE 5004

CMD ["./nodo_tai"]