FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.* ./

RUN go mod download

COPY . .

RUN swag init -g ./main.go ./adapter/input/controller/user.controller.go -o ./main.go
RUN go build -o main main.go

EXPOSE 8080

CMD ["./main"]
