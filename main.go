package main

import (
	"fmt"
	"net/http"
	"github.com/test/restApi/router"
	"github.com/test/restApi/log"
)

func init() {
	err := log.LogInit()
	if err != nil {
		fmt.Println("panic: log init error")
	}
}

func main ()  {
	router := router.NewRouter()
	log.Info("server is running...")
	fmt.Println("server is running...")
	log.ERROR.Fatal(http.ListenAndServe(":8000", router))
}
