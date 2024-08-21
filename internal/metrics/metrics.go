package metrics17

import (
	"expvar"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/felixge/httpsnoop"
)

var version = "1.0.0"

func Metrics(next http.HandlerFunc) http.HandlerFunc {
	expvar.NewString("version").Set(version)
	expvar.Publish("goroutines", expvar.Func(func() interface{} {
		return runtime.NumGoroutine()
	}))

	expvar.Publish("timestamp", expvar.Func(func() interface{} {
		return time.Now().Unix()
	}))
	totalRequest := expvar.NewInt("total_request_received")
	totalResponse := expvar.NewInt("total_response_sent")
	totalProcessingTime := expvar.NewInt("total_processing_time_us")
	totalResponsesByStatus := expvar.NewMap("total_responses_sent_by_status")

	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		totalRequest.Add(1)

		metrics := httpsnoop.CaptureMetrics(next, w, r)

		duration := time.Since(start).Microseconds()
		totalProcessingTime.Add(duration)
		totalResponse.Add(1)
		totalResponsesByStatus.Add(strconv.Itoa(metrics.Code), 1)
	}
}
