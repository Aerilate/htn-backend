FROM golang:1.18 AS builder
WORKDIR /src/

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=1 go build -a -ldflags '-linkmode external -extldflags "-static"' -o app


FROM alpine
WORKDIR /app/

COPY --from=builder /src/app ./
COPY db/migrate ./db/migrate
COPY *.json ./
CMD ["./app", "-i"]
