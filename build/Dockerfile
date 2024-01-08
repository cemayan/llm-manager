FROM golang:1.21.5-alpine3.19 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/git-observer  cmd/main.go

FROM scratch
COPY --from=builder /app/git-observer /app/git-observer
EXPOSE 8996
ENTRYPOINT ["/app/git-observer"]