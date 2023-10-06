FROM golang

WORKDIR /crypto-app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build  -C ./cmd -o /crypto-server

CMD ["../crypto-server",  "-docker"]