package middlewares

import (
	"backend/models"
	"backend/repository/datastore"
	"crypto/tls"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis_rate/v10"
	"github.com/gorilla/schema"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func EnvMiddleware(zapLogger *zap.Logger, validate *validator.Validate, decoder *schema.Decoder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		postgeSQLConn, rateLimiter, redisClient := initialiseDatastores()
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient := &http.Client{Transport: tr}
		env := &models.Env{
			Logger:         zapLogger,
			AuthDtos:       &models.Auth{},
			DbConn:         postgeSQLConn,
			RedisClient:    redisClient,
			HttpClient:     httpClient,
			RateLimiter:    rateLimiter,
			Ctx:            ctx,
			Validator:      validate,
			Decoder:        decoder,
			RequestContext: make(map[string]interface{}),
		}
		ctx.Set("env", env)
		ctx.Next()
		ctx.Set("env", nil)
		env.DbConn = nil
		env.RedisClient = nil
		env.HttpClient = nil
		env.RateLimiter = nil
		env.Validator = nil
		env.Decoder = nil
		env.RequestContext = nil
	}
}

func initialiseDatastores() (*gorm.DB, *redis_rate.Limiter, *redis.Client) {
	datastore.Get()
	return datastore.PostgeSQLConn, datastore.RateLimiter, datastore.RedisClient
}
