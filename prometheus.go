package common

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func NewPrometheus(add string) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(add, nil)
		if err != nil {
			fmt.Println("prometheus启动失败", err)
		}
	}()
}
