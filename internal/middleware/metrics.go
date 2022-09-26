package middleware

import (
	metrics "github.com/yalagtyarzh/online-chat/internal/metrics_kit"
	"github.com/yalagtyarzh/online-chat/pkg/utils"
	"net/http"
	"time"
)

func MetricsMiddleware(next http.Handler, m *metrics.HTTPMetrics) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		m.TotalRequests().Inc()
		m.CurrentConnections().Inc()
		defer m.CurrentConnections().Dec()

		rsi := NewResponseStatusInterceptor(w)
		next.ServeHTTP(rsi, r)

		if utils.IsBetween(rsi.statusCode, 400, 499) {
			m.ClientErrorRequests().Inc()
		} else if utils.IsBetween(rsi.statusCode, 500, 599) {
			m.ServerErrorRequests().Inc()
		} else {
			m.OKRequests().Inc()
			m.OKRequestTimer().Observe(time.Since(startTime).Seconds())
		}
	})
}
