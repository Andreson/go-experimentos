package boostrap

import (
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	APP_CONFIG         = LoadEnvApp()
	REDIS_CONFIG       = LoadEnvRedis()
	WEB_HANDLER_CONFIG = LoadEnvWebHandler()
)
var Log *log.Entry

func init() {
	LogConfig()
}

func LoadEnvWebHandler() WebHandlerConfig {
	var web WebHandlerConfig

	if invalidCacheKey := os.Getenv("SDC_INVALID_CACHE_KEY"); invalidCacheKey == "" {
		web.InvalidCachKey = "x-clean-cache"
	} else {
		web.InvalidCachKey = invalidCacheKey
	}

	if invalidCacheKey := os.Getenv("SDC_ID_CALLER_CLIENT"); invalidCacheKey == "" {
		web.IdCallerLog = "x-trace-id"
	} else {
		web.IdCallerLog = invalidCacheKey
	}

	if contentType := os.Getenv("SDC_CONTENT_TYPE"); contentType == "" {
		web.ContentTypeResponse = "application/json"
	} else {
		web.ContentTypeResponse = contentType
	}
	log.Info("Configurações WebHandler :  ", web)
	return web
}

func LoadEnvApp() AppConfig {
	var app AppConfig
	var err error

	app.Port = loadEnvString("SDC_SERVER_PORT", "9200")

	app.PrometheusEnable = loadEnvBool("SDC_ENABLE_PROMETHEUS", "false")

	app.RetryConnProvider = loadEnvInt("SDC_RETRY_CONNECTION", 5)

	app.PathParamIndex, _ = loadPathParamIndex(strings.Split(os.Getenv("SDC_PATHPARAM_INDEX"), ","))

	app.Enable, err = strconv.ParseBool(loadEnvString("SDC_ENABLE", "true"))

	app.LogLevel, err = log.ParseLevel(loadEnvString("SDC_LOG_LEVEL", "info"))
	if err != nil {
		log.Errorf("Erro ao carregar level log configurado %s |Carregando log level  default como info", err)
		app.LogLevel, _ = log.ParseLevel("info")
	}

	app.BussinesContainerAddr = loadEnvString("SDC_BUSSINES_URL", "http://localhost:8080")

	app.Provider = loadEnvString("SDC_CACHE_PROVIDER", "REDIS")
	app.Version = "1.2.5.1"
	log.Info("Configurações application  :  ", app)
	return app

}

func LoadEnvRedis() RedisConfig {
	var redis RedisConfig
	var err error
	redis.Port = loadEnvString("SDC_REDIS_PORT", "6379")

	redis.Host = loadEnvString("SDC_REDIS_HOST", "localhost")
	//carrega o tempo de vida do cache em minutos
	redis.Ttl = loadEnvInt("SDC_REDIS_TTL_MIN", 45)

	redis.PoolSize = loadEnvInt("SDC_REDIS_POOLSIZE", 15)

	redis.MaxConnAge = loadEnvInt("SDC_REDIS_MAXCONNAGE_SEC", 10)

	redis.IdleTimeout = loadEnvInt("SDC_REDIS_IDLETIMEOUT_SEC", 10)

	redis.WriteTimeout = loadEnvInt("SDC_REDIS_WRITETIMEOUT_SEC", 2)

	redis.ReadTimeout = loadEnvInt("SDC_REDIS_READTIMEOUT_SEC", 2)

	redis.DialTimeout = loadEnvInt("SDC_REDIS_DIALTIMEOUT_SEC", 1)

	redis.RetryConnection = loadEnvInt("SDC_REDIS_RETRYCONNECTION", 5)

	redis.User = os.Getenv("SDC_REDIS_USER")

	redis.Context = os.Getenv("SDC_REDIS_CONTEXT")

	redis.Db = loadEnvInt("SDC_REDIS_DB", 0)

	redis.Password = os.Getenv("SDC_REDIS_PASSWORD")

	redis.FailFast, err = strconv.ParseBool(loadEnvString("SDC_REDIS_FAIL_FAST", "false"))

	if err != nil { //setar "false" define  que app ira iniciar mesmo se conexao com provider falhar
		// apenas como poxy, sem cache
		redis.FailFast = false
	}

	log.Infof("Configurações Redis [%s] [%s] [%s] :  ", redis.Host, redis.FailFast, redis.Ttl)
	return redis
}
