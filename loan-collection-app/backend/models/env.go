package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis_rate/v10"
	"github.com/gorilla/schema"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Env struct {
	Logger         *zap.Logger
	AuthDtos       *Auth
	DbConn         *gorm.DB
	RedisClient    *redis.Client
	HttpClient     *http.Client
	RateLimiter    *redis_rate.Limiter
	Ctx            *gin.Context
	Validator      *validator.Validate
	Decoder        *schema.Decoder
	RequestContext map[string]interface{}
	PermissionList []string
}

type Auth struct {
	User *User
}
