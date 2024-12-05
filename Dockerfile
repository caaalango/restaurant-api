FROM golang:1.23 

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@v1.21.2
RUN CGO_ENABLED=0 go install github.com/cosmtrek/air@v1.49.0
RUN CGO_ENABLED=0 go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2

COPY go.mod go.sum ./
COPY air.toml ./

RUN go mod download
RUN go mod verify

EXPOSE 8080
EXPOSE 40000

ENTRYPOINT ["air", "-c", "air.toml"]
