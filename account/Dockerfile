FROM golang:1.21.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

COPY . ./

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -o /account-server

CMD ["../account-server",  "-docker"]