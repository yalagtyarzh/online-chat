package metrics

import "github.com/prometheus/client_golang/prometheus"

func NewCounter(namespace, name, help string) (prometheus.Counter, error) {
	counter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      name,
			Help:      help,
		})

	if err := prometheus.Register(counter); err != nil {
		return nil, err
	}

	return counter, nil
}

func NewGauge(namespace, name, help string) (prometheus.Gauge, error) {
	gauge := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      name,
			Help:      help,
		})

	if err := prometheus.Register(gauge); err != nil {
		return nil, err
	}

	return gauge, nil
}

func NewHistogram(namespace, name, help string) (prometheus.Histogram, error) {
	gauge := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      name,
			Help:      help,
		})

	if err := prometheus.Register(gauge); err != nil {
		return nil, err
	}

	return gauge, nil
}
