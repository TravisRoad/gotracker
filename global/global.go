package global

import (
	"github.com/TravisRoad/gotracker/config"
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config         config.Config
	DB             *gorm.DB
	CasbinEnforcer *casbin.Enforcer
	Viper          *viper.Viper
	Logger         *zap.Logger
	Redis          *redis.Client
	Mode           string // "DEBUG" "TEST" "PROD"
	WorkspacePath  string
)
