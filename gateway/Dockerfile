FROM golang

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

EXPOSE 8282

RUN CGO_ENABLED=0 GOOS=linux go build  -C ./cmd -o /gateway-server

CMD ["../gateway-server"]