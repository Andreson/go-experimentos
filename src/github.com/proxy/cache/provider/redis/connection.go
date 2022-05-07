package redis

import (
	"context"
	"errors"
	"time"

	boot "br.com.hdi.cross.cache/boostrap"
	"github.com/go-redis/redis/v8"
)

var (
	ErrNil    = errors.New("Nenhum registro encontrado no cache!")
	Ctx       = context.TODO()
	RedisAddr = boot.REDIS_CONFIG.Host + ":" + boot.REDIS_CONFIG.Port
)

//abre conexão com redis com  usuario e senha
func ConnAuth(address string) (client *redis.Client, err error) {

	if !boot.APP_CONFIG.Enable {
		boot.Log.Warn("Sidecar de cache foi desabilitado em suas configurações de ambiente. Suas requisições nao serão cacheadas")
		return nil, errors.New("Sidecar de cache foi desabilitado em suas configurações de ambiente. Suas requisições nao serão cacheadas")
	}

	boot.Log.Info("Iniciando conexão com Redis Memorystore")

	client = redis.NewClient(&redis.Options{
		Addr:        address,
		Username:    boot.REDIS_CONFIG.User,
		Password:    boot.REDIS_CONFIG.Password,
		DB:          boot.REDIS_CONFIG.Db,
		PoolSize:    boot.REDIS_CONFIG.PoolSize,
		MaxConnAge:  time.Second * time.Duration(boot.REDIS_CONFIG.MaxConnAge),
		IdleTimeout: time.Second * time.Duration(boot.REDIS_CONFIG.IdleTimeout),
		DialTimeout: time.Second * time.Duration(boot.REDIS_CONFIG.DialTimeout),
	})

	_, err = client.Ping(Ctx).Result()

	if err != nil {
		boot.Log.Error("--------- Erro PING redis ", err.Error())
		client.Close()
		return nil, err
	}

	return
}
