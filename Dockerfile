# Build stage
FROM golang:tip-alpine3.22 AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod ./

# Download dependencies (if any)
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o taxcal .

# Final stage
FROM alpine:latest

# Install ca-certificates for any HTTPS requests
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/taxcal .

# Copy the tax rates file
COPY --from=builder /app/tax_rates.json .

# Run the application
CMD ["./taxcal"]
