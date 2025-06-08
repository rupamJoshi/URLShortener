FROM golang:1.21

WORKDIR /app

COPY . .
RUN go mod download



RUN pwd
RUN ls

RUN CGO_ENABLED=0 GOOS=linux go build -o /short_url ./cmd/main.go

EXPOSE 8080
CMD ["/short_url"]