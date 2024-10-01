FROM golang:1.23.1-bullseye as builder

WORKDIR /app


COPY go.mod go.sum ./


RUN go version

RUN go mod tidy

COPY . .

RUN make build

FROM scratch
WORKDIR /app

COPY --from=builder /app/bin .
COPY --from=builder /app/assets ./assets

EXPOSE 3000

CMD ["/app/app"]

