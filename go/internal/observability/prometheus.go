package observability

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MountPrometheus は /metrics エンドポイントを登録します。
// Go/プロセスのデフォルトメトリクスのみを公開します（最小構成）。
func MountPrometheus(r *gin.Engine, path string) {
	if path == "" {
		path = "/metrics"
	}

	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	h := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	r.GET(path, func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) })
}
