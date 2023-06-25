package driver

import (
	"bytes"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"
	"time"
)

// TaskController controls task creation and their worker interaction.
type TaskController struct {
	tasks             chan task       //queue for sending tasks
	registeredWorkers map[string]bool //set of available registeredWorkers
	workersQueue      chan string     //queue of available workers
	regMutex          *sync.Mutex     //mutex for registration handle
	doneMutex         *sync.Mutex     //mutex for done handle
	sentTasks         map[string]task //map of sentTasks
	hc                http.Client     //http client for driver interaction with workers
	M                 int             //The number N of map tasks
	N                 int             //The number M of reduce tasks
	files             []string        //Slice of files sent to worker
}

// NewController is creating new TaskController instance
func NewController(M int, N int, c Conf) *TaskController {
	return &TaskController{
		regMutex:          new(sync.Mutex),
		doneMutex:         new(sync.Mutex),
		sentTasks:         make(map[string]task),
		registeredWorkers: make(map[string]bool),
		M:                 M,
		N:                 N,
		tasks:             make(chan task, N),
		files:             getFilenames(c.Inputs),
	}
}

// getFilenames gets full paths to files for controller
func getFilenames(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	result := make([]string, len(entries))
	for i, e := range entries {
		result[i] = path.Join(dir, e.Name())
	}
	return result
}

// Run is main function for task operation with workers
// waits for some time for all possible workers could register in driver
func (c *TaskController) Run() {
	c.createMapTasks()
	log.Println("Awaiting for workers")
	time.Sleep(20 * time.Second)
	if len(c.registeredWorkers) == 0 {
		log.Println("No workers registered, stopping driver")
		return
	}
	c.workersQueue = make(chan string, len(c.registeredWorkers))
	for k, _ := range c.registeredWorkers {
		c.workersQueue <- k
	}
	log.Printf("Sending %d map tasks\n", len(c.tasks))
	c.processTasks(true)
	log.Println("Done map tasks")
	c.createReduceTasks()
	log.Printf("Sending %d reduce tasks\n", len(c.tasks))
	c.processTasks(false)
	log.Println("Done reduce tasks")
	log.Println("All tasks done")
	log.Println("You may close this process")
	log.Println("This process will be automatically closed in 10 seconds")
	time.Sleep(10 * time.Second)
	log.Fatalln("Stopping driver")
}

// processTasks is main loop for sending tasks to workers
func (c *TaskController) processTasks(isMap bool) {
	var fun = func(t task, url string) {}
	var message string
	if isMap {
		fun = c.sendMapTask
		message = "All map tasks sent"
	} else {
		fun = c.sendReduceTask
		message = "All reduce tasks sent"
	}
	for {
		t, _ := <-c.tasks
		worker, _ := <-c.workersQueue
		fun(t, worker)
		if len(c.tasks) == 0 {
			log.Println(message)
			for {
				if len(c.sentTasks) == 0 {
					break
				}
			}
			if len(c.tasks) == 0 {
				break
			}
		}
	}
}

// sendMapTask sends map task to worker
func (c *TaskController) sendMapTask(t task, url string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for i, filename := range t.files {
		part, _ := writer.CreateFormFile("file"+strconv.Itoa(i), filename)
		data, _ := os.ReadFile(filename)
		part.Write(data)
	}
	writer.Close()
	req, err := http.NewRequest(
		"POST",
		url+"/map",
		body)
	if err != nil {
		return
	}
	req.Header.Set("M", strconv.Itoa(c.M))
	req.Header.Set("Taskid", t.id)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	c.sentTasks[url] = t
	do, err := c.hc.Do(req)
	if err != nil {
		c.tasks <- c.sentTasks[url]
		delete(c.sentTasks, url)
		return
	}
	if do.StatusCode != http.StatusOK {
		log.Println(do.StatusCode)
	}
}

// sendReduceTask sends reduce task to worker
func (c *TaskController) sendReduceTask(t task, url string) {
	req, err := http.NewRequest(
		"POST",
		url+"/reduce",
		nil)
	if err != nil {
		return
	}
	req.Header.Set("Taskid", t.id)
	c.sentTasks[url] = t
	do, err := c.hc.Do(req)
	if err != nil {
		c.tasks <- c.sentTasks[url]
		delete(c.sentTasks, url)
		return
	}
	if do.StatusCode != http.StatusOK {
		log.Println(do.StatusCode)
	}
}

// createMapTasks creates map tasks and putting it into tasks queue
func (c *TaskController) createMapTasks() {
	l := len(c.files)
	batchSize := int(math.Ceil(float64(l) / float64(c.N)))
	log.Printf("Creating map tasks, batch size = %d", batchSize)
	for i := 0; i < c.N && i <= l; i += batchSize {
		if len(c.files[i:]) <= batchSize {
			c.tasks <- task{
				strconv.Itoa(i), c.files[i:],
			}
		} else {
			c.tasks <- task{
				strconv.Itoa(i), c.files[i : i+batchSize],
			}
		}

	}
	log.Println("Map tasks created")
}

// createReduceTasks creates reduce tasks and putting it into tasks queue
func (c *TaskController) createReduceTasks() {
	log.Println("Creating reduce tasks")
	for i := 0; i < c.M; i++ {
		c.tasks <- task{
			strconv.Itoa(i), nil,
		}
	}
	log.Println("Reduce tasks created")
}
