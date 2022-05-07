package main

import (
	"net/http"

	"br.com.hdi.cross.cache/boostrap"
	boot "br.com.hdi.cross.cache/boostrap"
	"br.com.hdi.cross.cache/web"
)

func init() {

}

func main() {

	boot.Log.Infof("############  INTEGRATION CACHE SIDECAR  [ v%s ]  #################", boot.APP_CONFIG.Version)
	boot.Log.Info("Iniciando servidor na porta " + boostrap.APP_CONFIG.Port)

	web.NewGetEnpoint()

	if boostrap.APP_CONFIG.PrometheusEnable {
		web.Monitoring()
	}

	boot.Log.Infof("Cache Sidecar  em execução na porta %s", boostrap.APP_CONFIG.Port)

	boot.Log.Fatal(http.ListenAndServe(":"+boostrap.APP_CONFIG.Port, web.GetRouter))

}
