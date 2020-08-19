package global

import (
	"github.com/casbin/casbin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/zhenghuajing/demo/config"
	"go.uber.org/zap"
)

var (
	DB        *gorm.DB
	Enforcer  *casbin.Enforcer
	RedisPool *redis.Pool
	Log       *zap.SugaredLogger
	Config    *config.Config
	Viper     *viper.Viper
)
