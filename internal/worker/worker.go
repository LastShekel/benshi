package worker

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Main() {
	c := LoadConfig()
	if _, err := os.Stat(c.Files); os.IsNotExist(err) {
		err = os.Mkdir(c.Files, 0777)
		if err != nil {
			panic("Failed to create files folder")
		}
	}
	c.WorkerPort = getAvailablePort(c.WorkerPort)
	r := NewRouter(c)
	go register("http://127.0.0.1:"+c.DriverPort+"/register", c.WorkerPort)

	log.Fatal(http.ListenAndServe(":"+c.WorkerPort, r))
}

// register sends register request to driver until success
func register(url string, port string) {
	values := map[string]string{"url": fmt.Sprintf("http://127.0.0.1:%s", port)}
	for {
		post, err := SendJson(url, values)
		if err != nil {
			fmt.Println("Failed to register in driver")
			time.Sleep(5 * time.Second)
			continue
		}
		if post.StatusCode == http.StatusCreated {
			fmt.Println("Worker registered successfully")
			break
		}
	}
}

// getAvailablePort looking for next available port to create new worker
func getAvailablePort(port string) string {
	ln, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't listen on port %q: %s\n", port, err)
		atoi, err := strconv.Atoi(port)
		if err != nil {
			panic("Available ports not found")
		}
		return getAvailablePort(strconv.Itoa(atoi + 1))
	}

	_ = ln.Close()
	fmt.Printf("TCP Port %q is available\n", port)
	return port
}
