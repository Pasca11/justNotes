package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	LoginCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "login_count",
		Help: "Number of login attempts",
	})
	RegisterCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "Register_counter",
		Help: "Counts /register uses",
	})
	RequestHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "RequestHistogram",
		Help:    "Histogram of HTTP request latencies",
		Buckets: prometheus.DefBuckets,
	},
		[]string{"method", "path"})
)

func init() {
	prometheus.MustRegister(LoginCount)
	prometheus.MustRegister(RegisterCounter)
	prometheus.MustRegister(RequestHistogram)
}
