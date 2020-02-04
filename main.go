package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
)

//
//import (
//	"log"
//	"net/http"
//
//	"github.com/prometheus/client_golang/prometheus"
//	"github.com/prometheus/client_golang/prometheus/promhttp"
//)
//
//var (
//	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
//		Name: "cpu_temperature_celsius",
//		Help: "Current temperature of the CPU.",
//	})
//	hdFailures = prometheus.NewCounterVec(
//		prometheus.CounterOpts{
//			Name: "hd_errors_total",
//			Help: "Number of hard-disk errors.",
//		},
//		[]string{"device"},
//	)
//)
//
//func init() {
//	// Metrics have to be registered to be exposed:
//	prometheus.MustRegister(cpuTemp)
//	prometheus.MustRegister(hdFailures)
//}
//
//func main() {
//	cpuTemp.Set(65.3)
//	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
//
//	// The Handler function provides a default handler to expose metrics
//	// via an HTTP server. "/metrics" is the usual endpoint for that.
//	http.Handle("/metrics", promhttp.Handler())
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}
func main() {
	//go func() {
	//	http.ListenAndServe(":8088", nil)
	//}()
	//for i := 0; i < 100; i++ {
	//	time.Sleep(time.Second)
	//	fmt.Println(":", i)
	//}
	//ch := make(chan string)
	//<-ch

	var count int32 = 0
	server := &http.Server{
		Addr: ":4444",
		Handler: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Set("Connection", "close")
		}),
		ConnContext: func(ctx context.Context, c net.Conn) context.Context {
			atomic.AddInt32(&count, 1)
			if c2 := ctx.Value("count"); c2 != nil {
				fmt.Printf("发现了遗留数据: %d\n", c2.(int32))
			}
			fmt.Printf("本次数据: %d\n", count)
			return context.WithValue(ctx, "count", count)
		},
	}
	go func() {
		panic(server.ListenAndServe())
	}()

	var err error

	fmt.Println("第一次请求")
	_, err = http.Get("http://localhost:4444/")
	if err != nil {
		panic(err)
	}
	fmt.Println("\n第二次请求")

	_, err = http.Get("http://localhost:4444/")
	if err != nil {
		panic(err)
	}
}
