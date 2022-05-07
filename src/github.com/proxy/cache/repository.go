package cache

import (
	"errors"

	boot "br.com.hdi.cross.cache/boostrap"
	"br.com.hdi.cross.cache/cache/domain"
)

var RepositoryCache domain.CacheRepository

func init() {
	boot.Log.Info(" --------------------------  INSTANCIANDO  PROVIDER CACHE --------------------------")
	RepositoryCache = FactoryProvider()
	RepositoryCache.Connect()

}

//Faz a busca de dados no redis, recebe os parametros do request, e executa a func que gera a key no cache
func FindDataCache(parametersRequest string) domain.CacheResponse {
	var cacheResponse domain.CacheResponse

	if !boot.APP_CONFIG.Enable { //valido se o sidecar subiu desabilitado
		boot.Log.Warn("SIDECAR DESABILITADO PARA REQUESTS")
		return cacheResponse
	}

	if !tryConnectionIsError() { //falhou ao tentar conectar no provider
		cacheResponse.SetStatus(errors.New("Erro ao conectar no provider de cache "))
		return cacheResponse
	}
	key := GenerateCacheKey(parametersRequest)
	cacheResponse = RepositoryCache.Get(key)
	if cacheResponse.Data() == "" {
		boot.Log.Debugf("Dados nao encontrados cache para    key [ %s ]  | params [ %s ]  | MSG  [ %s ] ", key, parametersRequest, cacheResponse.Error())
		return cacheResponse
	} else {
		boot.Log.Debugf("Dados retornados do cache :  %s ", cacheResponse.Data())
		boot.Log.Infof("Resultados do  cache para : key [ %s ]  | params [ %s ]  ", key, parametersRequest)
		return cacheResponse
	}
}

//Adiciona os dados  no redis, recebe os parametros do request, e executa a func que gera a key no cache
func SetDataCache(parametersRequest string, data string) error {

	key := GenerateCacheKey(parametersRequest)

	if !tryConnectionIsError() { //falhou ao tentar conectar no provider
		return errors.New("Erro ao conectar no provider de cache")
	}
	boot.Log.Debugf(" Add dados no  cache client-request: [%s]  :  key [ %s ] | params [ %s ] : ", key, parametersRequest)
	status := RepositoryCache.Set(key, data)
	return status.Status.Err
}

//Deleta os dados  no redis, recebe os parametros do request, e executa a func que gera a key no cache
func DelDataCache(parametersRequest string) error {
	key := GenerateCacheKey(parametersRequest)

	status := RepositoryCache.Del(key)
	if status.Error() != "" {
		boot.Log.Error("******* Erro invalidar cache **************")
	}
	boot.Log.Infof(" Del cache  key [ %s ] | params [ %s ] | deletado ?  [ %t ]  ", key, parametersRequest, status.Error() == "")
	return status.Status.Err
}
