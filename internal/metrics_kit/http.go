package metrics_kit

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/yalagtyarzh/online-chat/pkg/metrics"
)

type HTTPMetrics struct {
	totalRequests       prometheus.Counter
	okRequests          prometheus.Counter
	serverErrorRequests prometheus.Counter
	clientErrorRequests prometheus.Counter
	currentConnections  prometheus.Gauge
	okRequestTimer      prometheus.Histogram
}

func (H *HTTPMetrics) TotalRequests() prometheus.Counter {
	return H.totalRequests
}

func (H *HTTPMetrics) OKRequests() prometheus.Counter {
	return H.okRequests
}

func (H *HTTPMetrics) ServerErrorRequests() prometheus.Counter {
	return H.serverErrorRequests
}

func (H *HTTPMetrics) ClientErrorRequests() prometheus.Counter {
	return H.clientErrorRequests
}

func (H *HTTPMetrics) CurrentConnections() prometheus.Gauge {
	return H.currentConnections
}

func (H *HTTPMetrics) OKRequestTimer() prometheus.Histogram {
	return H.okRequestTimer
}

func NewHTTPMetrics() (*HTTPMetrics, error) {
	totalRequests, err := metrics.NewCounter(namespace, "nb_req", "Total amount of HTTP requests")
	if err != nil {
		return nil, err
	}

	// Init metric for amount of OK HTTP requests
	okRequests, err := metrics.NewCounter(namespace, "nb_req_2xx", "Amount of OK HTTP requests")
	if err != nil {
		return nil, err
	}

	// Init metric for amount of server error HTTP requests
	serverErrorRequests, err := metrics.NewCounter(namespace, "nb_req_5xx",
		"Amount of server error HTTP requests")
	if err != nil {
		return nil, err
	}

	// Init metric for amount of client error HTTP requests
	clientErrorRequests, err := metrics.NewCounter(namespace, "nb_req_4xx",
		"Amount of client error HTTP requests")
	if err != nil {
		return nil, err
	}

	// Init metric for current amount of TCP connections
	currentConnections, err := metrics.NewGauge(namespace, "http_current_connections_total", "Current amount of TCP connections")
	if err != nil {
		return nil, err
	}

	// Init metric for OK request processing
	okRequestTimer, err := metrics.NewHistogram(namespace, "req_time",
		"Processing time for a OK request")
	if err != nil {
		return nil, err
	}

	return &HTTPMetrics{
		totalRequests:       totalRequests,
		okRequests:          okRequests,
		serverErrorRequests: serverErrorRequests,
		clientErrorRequests: clientErrorRequests,
		currentConnections:  currentConnections,
		okRequestTimer:      okRequestTimer,
	}, nil
}
