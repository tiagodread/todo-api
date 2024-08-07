FROM golang:1.22

WORKDIR /app

COPY . /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o todo-api cmd/main.go

WORKDIR /app

EXPOSE 8080

RUN chmod +x todo-api

CMD [ "./todo-api" ]