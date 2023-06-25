package driver

import (
	"log"
	"net/http"
)

type task struct {
	id    string
	files []string
}

// Main worker process is running
func Main(N, M int) {
	log.Println("Driver started")
	c := LoadConfig()
	log.Println("Config loaded")
	controller := NewController(M, N, c)
	r := NewRouter(controller)
	go controller.Run()
	log.Println("Driver is ready!")
	log.Fatal(http.ListenAndServe(":"+c.DriverPort, r))
}
