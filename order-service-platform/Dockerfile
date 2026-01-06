FROM golang:1.25.1-alpine

WORKDIR /app

# 1️⃣ 複製 go mod
COPY go.mod go.sum ./
RUN go mod download

# 2️⃣ 複製全部原始碼
COPY . .

# 3️⃣ 編譯 api-gateway
RUN go build -o api-gateway ./service/api-gateway

EXPOSE 8080

CMD ["./api-gateway"]
