FROM golang

WORKDIR /account-app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -o /account-server

CMD ["../account-server",  "-docker"]