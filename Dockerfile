FROM golang:1.21.1

WORKDIR /app

COPY go.mod go.sum ./

RUN apt-get update && apt-get install -y gcc
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o /my_app ./cmd/main.go

EXPOSE 7540

ENV HTTP_HOST=0.0.0.0
ENV HTTP_PORT=7540
ENV TODO_DBFILE=./scheduler.db
ENV TODO_PASSWORD=123456
ENV LOGLEVEL=info

CMD ["/my_app"]
