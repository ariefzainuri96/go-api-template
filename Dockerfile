FROM golang:1.24-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o ./bin/ ./cmd/api/

EXPOSE 8080

CMD [ "air", "-c", ".air.toml" ]