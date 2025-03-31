package global

import (
	"database/sql"
	"go-backend-api/pkg/logger"
	"go-backend-api/pkg/setting"

	"github.com/casbin/casbin/v2"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config      setting.Config
	Logger      *logger.LoggerZap
	Mdb         *gorm.DB
	Rdb         *redis.Client
	Mdbc        *sql.DB
	Enforcer    *casbin.SyncedEnforcer
	MinioClient *minio.Client
)
