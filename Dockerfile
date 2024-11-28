FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY . .

RUN go clean --modcache
RUN GOOS=linux go build -ldflags="-w -s" -o main ./cmd/server/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 3333

CMD [ "./main" ]
