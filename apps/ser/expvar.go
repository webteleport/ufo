package ser

import (
	"expvar"
	"runtime"
	"time"
)

var memStats = expvar.NewInt("memory")

func collectMemstats() {
	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		memStats.Set(int64(m.Alloc))
		time.Sleep(time.Second) // Adjust the sleep duration as needed
	}
}
