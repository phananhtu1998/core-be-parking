# Bước 1: Build Go binary
FROM golang:alpine AS builder
WORKDIR /build

COPY . .  
RUN go mod download 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o crm.gobackend.com ./cmd/server

# Bước 2: Đóng gói vào scratch để giữ image nhẹ
FROM scratch
WORKDIR /root/

# Copy binary đã build
COPY --from=builder /build/crm.gobackend.com /

# Copy config
COPY ./config /root/config

ENTRYPOINT ["/crm.gobackend.com", "config/local.config"]
