package boostrap

import (
	log "github.com/sirupsen/logrus"
)

type AppConfig struct {
	Port                  string
	BussinesContainerAddr string // url para chamada do serviço principar que sera a origem dos dados cacheados
	LogLevel              log.Level
	Enable                bool // desabilita o uso do sidecar, e usado programaticamente em caso de falha na conexão com redis
	PrometheusEnable      bool
	Provider              string
	RetryConnProvider     int
	PathParamIndex        []int
	Version               string
}

type RedisConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	Context         string
	Db              int
	Ttl             int //TEMPO EM MINUTOS
	RetryConnection int //TEMPO EM MINUTOS
	PoolSize        int
	MaxConnAge      int //tempo em segundos
	IdleTimeout     int //tempo em segundos
	ReadTimeout     int
	WriteTimeout    int
	DialTimeout     int
	FailFast        bool
}

//struct para parametrizar ações na parte de manipua
type WebHandlerConfig struct {
	InvalidCachKey      string // nome do header que sera utilizado para invalidar o cache
	IdCallerLog         string // campo usado para debugs. nome do header que vai identificar meu client que fez a chamada  ao ws
	ContentTypeResponse string // contente type retornadno no serviço, default e application/json
}
