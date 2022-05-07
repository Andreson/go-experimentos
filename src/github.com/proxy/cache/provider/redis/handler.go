package redis

import (
	"time"

	boot "br.com.hdi.cross.cache/boostrap"
	"br.com.hdi.cross.cache/cache/domain"
)

func (rp *RedisProvider) Get(key string) domain.CacheResponse {
	val, err := rp.Client.Get(Ctx, key).Result()

	resp := domain.CacheResponse{}
	resp.SetStatus(err)

	resp.SetData(val)
	return resp

}
func (rp *RedisProvider) Set(key string, data string) domain.CacheResponse {

	statusCmd := rp.Client.Set(Ctx, key, data, time.Minute*time.Duration(boot.REDIS_CONFIG.Ttl)).Err()
	resp := domain.CacheResponse{}
	resp.SetStatus(statusCmd)

	if statusCmd != nil {
		boot.Log.Error("OCORREU UM ERRO NAO ESPERADO AO ADICIONAR DADO NO CACHE :  ", statusCmd)
		rp.ConneError = statusCmd
		return resp
	}

	return resp
}
func (rp *RedisProvider) Del(key string) domain.CacheResponse {

	error := rp.Client.Del(Ctx, key).Err()
	resp := domain.CacheResponse{}
	resp.SetStatus(error)

	if error != nil {
		rp.ConneError = error
		return resp
	}

	return domain.CacheResponse{}
}

func (rp *RedisProvider) Connect() domain.CacheConnection {

	rp.Client, rp.ConneError = ConnAuth(RedisAddr)

	if rp.ConneError != nil {
		boot.Log.Errorf("---------NAO FOI POSSIVEL CONNECTAR AO REDIS:  HOST %s PORT %d | ERROR : %S ", boot.REDIS_CONFIG.Host, boot.REDIS_CONFIG.Port)
	}
	return domain.CacheConnection{Err: rp.ConneError, Conn: rp.Client}
}

func (rp *RedisProvider) Error() error {

	return rp.ConneError
}
