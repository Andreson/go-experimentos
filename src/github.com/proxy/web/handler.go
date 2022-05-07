package web

//contem os metodos auxiliares para manipulação dos parametros
import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	boot "br.com.hdi.cross.cache/boostrap"
	"github.com/Jeffail/gabs"
)

//extrai os parametros do requests do path e da query ,e retorna uma string concatenada para usar como chave de cache no provider
func ExtractParamsKey(r *http.Request) (allParamns string) {

	pathParams := getPathParams(r)
	queryStringParams := getQueryStringParams(r)

	boot.Log.Debugf("-----------Parametros extraidos: path=> [%s]  query => [%s]   ", pathParams, queryStringParams)

	return pathParams + queryStringParams
}

//retorna uma string contendo os path params a serem utilizados na key do cache
func getPathParams(request *http.Request) (allParamns string) {
	params := strings.Split(request.URL.Path, "/")
	var concat ConcatString
	for _, index := range boot.APP_CONFIG.PathParamIndex {
		concat.Add(params[index])
	}
	value := concat.Build()

	return value

}

// retorna dos os parametros de query string de um request a serem utilizados na key do cache
func getQueryStringParams(request *http.Request, skipEndpointKey ...bool) (allParamns string) {

	var skipKey bool

	if skipEndpointKey == nil {
		skipKey = false
	} else {
		skipKey = skipEndpointKey[0]
	}

	params := request.URL.Query()
	var concat ConcatString
	keysOrder := orderKeysMap(params) //foi necessario ordenar os campos apos observar q o mesmo request feito varias vezes, recebe
	//os parametros em ordens diferentes, o que gera diferentes keys, para o mesmo conjunto de dados

	for _, key := range keysOrder {
		//ignorar key endpoints
		if key != "key" && !skipKey {
			boot.Log.Debugf(" Parametros encontrados:  key { %s} - value {%s}", key, params[key][0])
			concat.Add(key).Add("=").Add(params[key][0])
		}
	}

	return concat.Build()
}

//faz o cast do http response e retornar uma string no formato json
func ParseHttpResponseJson(reqApi *http.Response) (string, error) {

	if reqApi.Body == nil {
		return "", errors.New("Body request está vazio, nao e possivle converter objeto para json")
	}
	responseData, err := ioutil.ReadAll(reqApi.Body)
	if err != nil {
		boot.Log.Error(" Erro ao ler resposta chamada webservice: ", err)
		return "", err
	}
	jsonParsed, err := gabs.ParseJSON(responseData)

	if err != nil {
		boot.Log.Error("  ParseHttpResponseJson Erro ao fazer parse response json : %s ", err)
		printError(responseData)
		return "", err
	}
	return jsonParsed.String(), nil
}

func printError(responseData []byte) {

	var obj json.RawMessage
	if err := json.Unmarshal(responseData, &obj); err != nil {
		boot.Log.Info("Erro parse JSon ", err.Error())
	}

	boot.Log.Infof("%+v\n", obj)
	boot.Log.Infof("JSON: %s\n", obj)

}

//Recebe uma string no formato json, e retorna um json no formato correto para resposta da API
func ParseStringToJson(jsonString string) (string, error) {

	jsonParsed, err := gabs.ParseJSON([]byte(jsonString))

	if err != nil {
		boot.Log.Error(" Erro ao fazer parse response json : ", err)
		return "", err
	}
	return jsonParsed.String(), nil
}

func orderKeysMap(values map[string][]string) []string {
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}
