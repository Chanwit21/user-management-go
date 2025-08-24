# build stage
FROM golang:1.21-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/beast

# final image
FROM scratch
COPY --from=builder /app/beast /beast
EXPOSE 3000
ENTRYPOINT ["/beast"]
