package web

import (
	"encoding/json"
	"net/http"

	"br.com.hdi.cross.cache/boostrap"

	"br.com.hdi.cross.cache/cache"
)

// criar um endpoint com o verbo http GET
func HealthEnpoint() {

	GetRouter.Path("/health").HandlerFunc(healthCacheGetData).Methods("GET")
	GetRouter.Path("/info").HandlerFunc(infoHeathCheck).Methods("GET")
	return
}

func infoHeathCheck(response http.ResponseWriter, request *http.Request) {
	var responseModel = HeathResponse{RedisConfiguration: boostrap.REDIS_CONFIG,
		AplicationConfiguration: boostrap.APP_CONFIG,
		WebConfiguration:        boostrap.WEB_HANDLER_CONFIG}
	keyHealthCache := "health"
	setResponse := cache.SetDataCache(keyHealthCache, "Let's Go!")
	getReponse := cache.FindDataCache(keyHealthCache)

	if getReponse.Status.Err != nil || setResponse != nil {
		if !boostrap.REDIS_CONFIG.FailFast {
			response.WriteHeader(210)
		} else {
			response.WriteHeader(503)
		}
		responseModel.Mensage = "Falha ao conectar  provider de cache "
	} else {
		responseModel.Mensage = "Conexao com provider executada com sucesso! Cross integration cache Sidecar em execução!"
		response.WriteHeader(200)
	}

	responseJson, err := json.Marshal(responseModel)
	if err != nil {
		boostrap.Log.Error("Erro ao deserializar resposta '/info' cache sidecar", err)
		return
	}
	response.Write(responseJson)
	response.Header().Set("Content-Type", "application/json")
}

func healthCacheGetData(response http.ResponseWriter, request *http.Request) {
	keyHealthCache := "health" // + strconv.FormatInt((time.Now().UnixNano()/1e6), 10)
	//cache.Operations.Connect()
	setResponse := cache.SetDataCache(keyHealthCache, "Let's Go!")
	getReponse := cache.FindDataCache(keyHealthCache)

	if getReponse.Status.Err != nil || setResponse != nil {

		boostrap.Log.Errorf("Falha ao conectar provider Redis. Executando como proxy, sem cache. [v%s]", boostrap.APP_CONFIG.Version)
		if !boostrap.REDIS_CONFIG.FailFast {
			response.WriteHeader(210)
			response.Write([]byte("Falha ao conectar provider Redis. Executando como proxy, sem cache."))
			return
		}
		response.WriteHeader(503)
		response.Write([]byte("Ocorreu um erro não esperado ao conectar no provider de cache!"))

		return
	} else {
		boostrap.Log.Debugf("Conexao com provider executada com sucesso! Cross integration cache Sidecar em execução!")
		response.WriteHeader(200)
	}

	response.Write([]byte(getReponse.Data()))

}

type HeathResponse struct {
	Mensage                 string
	RedisConfiguration      boostrap.RedisConfig
	AplicationConfiguration boostrap.AppConfig
	WebConfiguration        boostrap.WebHandlerConfig
}
