package main

import (
	"ExportData/src/common/setting"
	"ExportData/src/common/validator"
	"ExportData/src/routers"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"time"
)

func main() {
	binding.Validator = new(validator.DefaultValidator)
	router := routers.InitRouter()
	conf := setting.Config.Server
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", conf.Port),
		Handler:        router,
		ReadTimeout:    conf.ReadTimeout * time.Second,
		WriteTimeout:   conf.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}