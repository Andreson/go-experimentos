package cache

import (
	"crypto/sha1"
	"fmt"

	boot "br.com.hdi.cross.cache/boostrap"
	"br.com.hdi.cross.cache/cache/domain"
	"br.com.hdi.cross.cache/cache/provider/memory"
	"br.com.hdi.cross.cache/cache/provider/redis"
)

var retryCount int

//Gera a key do cache no redis, gera um sha1 como key usando os parametros recebidos
func GenerateCacheKey(paramns string) string {
	var hash = sha1.New()

	hash.Write([]byte(paramns))

	boot.Log.Debugf("Gerando chave de cache para parametros [ %s ] ", paramns)

	hashKey := fmt.Sprintf("%x", hash.Sum(nil))

	boot.Log.Debugf("-------------------------Chave gerada  [ %s ]", hashKey)
	hash.Reset()
	return hashKey
}

func FactoryProvider() domain.CacheRepository {
	boot.Log.Infof("Provider cache : %s ", boot.APP_CONFIG.Provider)
	switch provider := boot.APP_CONFIG.Provider; provider {
	case "MEMORY":
		boot.Log.Info("Instanciando sidecar cache memory")
		mem := &memory.MemoryProvider{}
		return mem

	case "REDIS":
		boot.Log.Info("Instanciando sidecar cache Memorystore Redis")
		redis := &redis.RedisProvider{}
		return redis
	default:
		boot.Log.Info("Instanciando sidecar cache memory")
		return &memory.MemoryProvider{}
	}
}

//Tentar  se conectar ao provider, e retorna false casa haja erro ao conectar no mesmo
func tryConnectionIsError() bool {

	if RepositoryCache.Error() != nil {

		if retryCount <= boot.APP_CONFIG.RetryConnProvider {
			boot.Log.Warnf("Erro ao connectar provider. Tentando nova conexao [%d de %d]", boot.APP_CONFIG.RetryConnProvider, retryCount)
			retryCount++
			if RepositoryCache.Connect().Err != nil { //retry
				return false
			}
		}
		return false
	}
	retryCount = 0
	return true
}
