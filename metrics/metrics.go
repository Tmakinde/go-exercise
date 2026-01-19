package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Count profile hits
	ProfileRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "profile_requests_total",
			Help: "Total number of profile endpoint hits",
		},
	)

	// Track login attempts by result
	LoginAttemptsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "login_attempts_total",
			Help: "Total login attempts by result",
		},
		[]string{"result"}, // success | failure
	)
)

// Init registers the metrics
func Init() {
	prometheus.MustRegister(ProfileRequestsTotal)
	prometheus.MustRegister(LoginAttemptsTotal)
}