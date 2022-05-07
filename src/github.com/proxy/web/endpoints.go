package web

//contem os metodos necessario  para criação dos endpoints http que ira receber os requests no pod
import (
	"net/http"

	boot "br.com.hdi.cross.cache/boostrap"
	"br.com.hdi.cross.cache/cache"
	"br.com.hdi.cross.cache/cache/domain"
	"github.com/gorilla/mux"
)

var GetRouter *mux.Router

//accoes pre-carregadas ao inicializar pacote
func init() {
	GetRouter = mux.NewRouter()
}

// criar um endpoint com o verbo http GET
func NewGetEnpoint() {
	HealthEnpoint()
	GetRouter.PathPrefix("/").HandlerFunc(findData).Methods("GET")
	return
}

func findData(response http.ResponseWriter, request *http.Request) {

	extractedParamnsCache := ExtractParamsKey(request)
	var responseData domain.CacheResponse

	response.Header().Set("Content-Type", boot.WEB_HANDLER_CONFIG.ContentTypeResponse)

	if request.Header.Get(boot.WEB_HANDLER_CONFIG.InvalidCachKey) == "true" { //invalida cache
		boot.Log.Infof("Invalidando request  para client  [%s] | endpoints key  [ %s ]", request.Header.Get(boot.WEB_HANDLER_CONFIG.IdCallerLog),
			request.URL.Query().Get("key"))
		_ = cache.DelDataCache(extractedParamnsCache)

	} else {
		responseData = cache.FindDataCache(extractedParamnsCache) //busca dados no redis
		if responseData.Data() != "" {
			response.Header().Set("x-cached", "true")
			response.Header().Set("x-cache-provider", boot.APP_CONFIG.Provider)
			response.Write([]byte(responseData.Data()))
		}
	}
	if responseData.Data() == "" {
		responseCall := findBussinesData(request, extractedParamnsCache)
		response.WriteHeader(responseCall.HttpResponse.StatusCode)
		response.Write([]byte(responseCall.DataContent))
	}

}

//Faz a chamada ao container de negocio caso os dados nao estejam no cache, e add a resposta do serviço no cache
func findBussinesData(request *http.Request, paramnsCache string) ResponseModel {
	boot.Log.Debugf("Dados nao encontrados no cache | client-traceId-request [ %s ] | paramns [ %s ] |", request.Header.Get(boot.WEB_HANDLER_CONFIG.IdCallerLog), paramnsCache)

	var cacheErr error
	responseModel := CallServiceMain(request) //chama container de negocio para buscar dados

	if len(responseModel.Error) > 0 {
		boot.Log.Error("Erro ao chamar serviço interno ", responseModel.Error)
	}

	if responseModel.HttpStatusCode >= 200 && responseModel.HttpStatusCode < 300 {

		boot.Log.Debugf("Cacheando dados | client-request [ %s ] | paramns [ %s ] |", request.Header.Get(boot.WEB_HANDLER_CONFIG.IdCallerLog), responseModel.DataContent)
		cacheErr = cache.SetDataCache(paramnsCache, responseModel.DataContent)
	}
	if cacheErr != nil {
		boot.Log.Errorf("Erro add dados no cache:  client-request [%s] | Error msg:   [ %s ] ",
			request.Header.Get(boot.WEB_HANDLER_CONFIG.IdCallerLog),
			cacheErr.Error())
	}
	return responseModel
}
