package global

import (
	"server/config"
	"server/utils/timer"
	"sync"

	ut "github.com/go-playground/universal-translator"

	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	DBList map[string]*gorm.DB
	REDIS  *redis.Client
	CONFIG config.Server
	VIP    *viper.Viper // 读取配置文件
	// LOG    *oplogging.Logger
	LOG                 *zap.Logger         // 日志
	Timer               timer.Timer         = timer.NewTimerTask()
	Concurrency_Control *singleflight.Group = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
	Validate   ut.Translator
	// 缓存
	// Cache config.Cache
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
