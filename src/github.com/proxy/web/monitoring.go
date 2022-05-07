package web

import (
	"net/http"

	boot "br.com.hdi.cross.cache/boostrap"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
)

const (
	metricsAddr = ":9210"
)

// criar um endpoint com o verbo http GET
func Monitoring() {

	boot.Log.Info("Configurando prometheus no sidecar de cache ")
	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})
	boot.Log.Info(mdlw)
	GetRouter.Use(std.HandlerProvider("/metrics", mdlw))

	go func() {
		boot.Log.Printf("Exposição de metricas habilitadas em    %s", metricsAddr)
		if err := http.ListenAndServe(metricsAddr, promhttp.Handler()); err != nil {
			boot.Log.Panicf("error while serving metrics: %s", err)
		}
	}()

	return
}
