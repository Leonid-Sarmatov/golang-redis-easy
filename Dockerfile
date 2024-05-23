FROM golang:1.22

WORKDIR /golang_redis

COPY . .

CMD ["go", "run", "Service/main.go"]