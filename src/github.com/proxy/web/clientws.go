package web

//define os metodos de chamadas para o serviço principal, container que contem os dados de negocio, ou chamada para um serviço externo
//caso tenha sido parametrizado para chamada externa
import (
	"net/http"

	boot "br.com.hdi.cross.cache/boostrap"
)

//chama o serviço principal que ira retornar os dados para o cache
func CallServiceMain(request *http.Request) ResponseModel {
	var responseModel ResponseModel
	resp, err := DoGetRequest(request)

	if err != nil {
		parseResponseData, err := ParseHttpResponseJson(resp)
		boot.Log.Error("Resposta nao esperada ao chamar webservice :  ", parseResponseData)
		responseModel.Error = append(responseModel.Error, err)
		return responseModel
	}

	responseModel = ResponseModel{HttpResponse: resp, HttpStatusCode: resp.StatusCode}

	parseResponseData, err := ParseHttpResponseJson(resp)
	responseModel.DataContent = parseResponseData

	return responseModel
}

//configura chamada GET e seta os heders da chamada original, executa a chamada e retorna o http response
func DoGetRequest(request *http.Request) (*http.Response, error) {
	var concat ConcatString

	urlRequest := concat.Add(boot.APP_CONFIG.BussinesContainerAddr).Add(request.URL.Path).Add("?").Add(request.URL.RawQuery).Build()
	boot.Log.Debugf("Chamando container de negocio:  [ %s ] ", urlRequest)
	req, err := http.NewRequest("GET", urlRequest, nil)

	if err != nil {
		boot.Log.Errorf("Erro criar novo request   [ %s ] ", err)
		return nil, err
	}
	boot.Log.Debugf("Headers retornados do container de negocio :  [ %s ] ", urlRequest)
	for key, value := range request.Header {
		req.Header.Set(key, value[0])
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		boot.Log.Errorf("Erro ao executar HTTP GET  |  \nHost [ %s ] |\n Body [ %s ] |\n Headers [ %s ] ", req.Host, req.Body, req.Header)
	} else {

		for key, value := range resp.Header {
			req.Header.Set(key, value[0])
			boot.Log.Debugf("\nHeaders retornados do container de negocio :  [ %s:%s ] ", key, value[0])
		}
	}

	return resp, err
}
