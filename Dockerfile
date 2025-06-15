# Stage 1: Build the Go binary with CGO enabled
FROM golangci/golangci-lint:v1.63 AS base
LABEL maintainer="hem.shrestha@ankaek.com"
LABEL last_modified="2024-09-01"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod tidy

# Stage 2: Lint
FROM base AS lint
RUN golangci-lint run --timeout 10m0s ./...

# Stage 3: Test
FROM base AS test
RUN go test -v -coverprofile=cover.out ./...
RUN go tool cover -func=cover.out

# Stage 4: Build
FROM base AS builder
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Stage 5: Final image with just the binary
FROM alpine:3.20 AS final
WORKDIR /app
RUN touch .env
COPY --from=builder /app/main .
COPY --from=builder /app/docs/* ./docs/
COPY --from=builder /app/templates/* ./templates/
COPY --from=builder /app/static/* ./static/
CMD ["/app/main"]
