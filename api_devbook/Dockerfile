FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build main.go

EXPOSE 8000

CMD ["./main", "-fake"]