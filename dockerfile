FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Para generar los go a partir de templ al momento de buildear
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate

# Compilamossss
RUN CGO_ENABLED=0 GOOS=linux go build -o blog-go ./cmd/main.go

# Runrun
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/blog-go .

COPY --from=builder /app/internal/view/public ./public
COPY --from=builder /app/internal/view/static ./static

EXPOSE 8080

ENTRYPOINT ["./blog-go"]

