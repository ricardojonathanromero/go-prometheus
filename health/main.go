package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"time"
)

func getEnv(k, v string) string {
	_v, ok := os.LookupEnv(k)
	if !ok {
		return v
	}

	return _v
}

func handleHealth(c *gin.Context) {
	c.String(http.StatusOK, "up")
}

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "health_app",
		Help: "The total number of processed events",
	})
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	recordMetrics()

	logger := log.New(os.Stdout, "health", 1)
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		logger.Fatalf("error setting trusted proxies: %v", err)
	}

	router.GET("/", handleHealth)
	router.GET("/metrics", prometheusHandler())
	port := getEnv("PORT", "8080")
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		logger.Fatal(err)
	}
}
