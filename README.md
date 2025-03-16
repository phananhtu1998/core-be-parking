# go-backend-api

cmd:

```
go mod init [tên thư mục chưa source]
```

Cài đặt GIN

Link 

```
https://github.com/gin-gonic/gin
```

CMD cài đặt:

```
go get -u github.com/gin-gonic/gin
```

Cài đặt zap để ghi log

```
go get -u go.uber.org/zap
```

Cài đặt Viper Config để cấu hình giá trị config

```
go get github.com/spf13/viper
```

import package test

```
go get github.com/stretchr/testify 
```

Thực thi lệnh này để chạy test chi tiết

```
go test -v
```

Tạo file kết quả test hiển thị lên HTML

```
cd [thư mục chứa file test vd: test/basic]
go test ./ -coverprofile=coverage.out
go tool cover -html=coverage -o coverage.html
```

Cài đặt lumberjack

```
go get github.com/natefinch/lumberjack
```

Cài đặt GORM để kết nối db

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

Cài đặt UUID

```
go get github.com/google/uuid
```

Cài đặt package redis

```
go get github.com/redis/go-redis/v9
```

Cài đặt wire go sử dụng dependence

```
go get github.com/google/wire/cmd/wire@latest
go install github.com/google/wire/cmd/wire@latest
```

Sau khi nhập các gói cần thiết phụ thuộc chạy lệnh để tạo file

```
cd internal/wire
wire
```

Migrate Schema GORM

```
go get -u gorm.io/gen
```

Cài đặt sqlc

```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Cài đặt goose

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

```
goose -dir sql/schema create pre_go_user_c sql
```

Tạo file để tạo bảng

```
makefile.bat create_migration 0001_pre_go_acc_user_verify_9999 
```

Lệnh này để tạo bảng

```
makefile.bat up_by_one 
```

Cài đặt swag

```
go install github.com/swaggo/swag/cmd/swag@latest
```

```
go get -u github.com/swaggo/swag
```

Cài đặt jwt

```
go get github.com/golang-jwt/jwt
```

Test độ chịu tải của server

```
artillery run load-test.yml
```

Lớp bảo vệ bên ngoài (global)

```
Resilience4j
```
