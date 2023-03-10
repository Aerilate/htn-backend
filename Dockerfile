FROM golang:1.18 AS builder
WORKDIR /src/

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=1 go build -a -ldflags '-linkmode external -extldflags "-static"' -o app
CMD ["sh"]


FROM alpine
WORKDIR /app/

COPY --from=builder /src/app ./
COPY migration migration/
COPY *.json ./

RUN ./app migrate
RUN ./app populate
CMD ["./app", "serve"]
