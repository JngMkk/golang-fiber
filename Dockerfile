FROM golang:1.21.4-alpine3.18 AS builder

RUN apk add --no-cache git

WORKDIR /build

# download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.2
RUN swag init

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o server .

# empty image
FROM scratch

COPY --from=builder [ "/build/server", "/build/.env", "/" ]

ENTRYPOINT [ "/server" ]