# From poggadaj-tcp of github.com/Oreeeee/poggadaj

# Build stage
FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o capstan

# Runtime stage
FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/capstan .

CMD ["./capstan"]