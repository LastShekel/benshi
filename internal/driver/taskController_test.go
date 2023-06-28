package driver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"sync"
	"testing"
)

func TestNewController(t *testing.T) {
	type args struct {
		M int
		N int
		c Conf
	}
	tests := []struct {
		name string
		args args
		want *TaskController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewController(tt.args.M, tt.args.N, tt.args.c), "NewController(%v, %v, %v)", tt.args.M, tt.args.N, tt.args.c)
		})
	}
}

func TestTaskController_Run(t *testing.T) {
	type fields struct {
		tasks             chan task
		registeredWorkers map[string]bool
		workersQueue      chan string
		regMutex          *sync.Mutex
		doneMutex         *sync.Mutex
		sentTasks         map[string]task
		hc                http.Client
		M                 int
		N                 int
		files             []string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TaskController{
				tasks:             tt.fields.tasks,
				registeredWorkers: tt.fields.registeredWorkers,
				workersQueue:      tt.fields.workersQueue,
				regMutex:          tt.fields.regMutex,
				doneMutex:         tt.fields.doneMutex,
				sentTasks:         tt.fields.sentTasks,
				hc:                tt.fields.hc,
				M:                 tt.fields.M,
				N:                 tt.fields.N,
				files:             tt.fields.files,
			}
			c.Run()
		})
	}
}

func TestTaskController_createMapTasks(t *testing.T) {
	type fields struct {
		tasks             chan task
		registeredWorkers map[string]bool
		workersQueue      chan string
		regMutex          *sync.Mutex
		doneMutex         *sync.Mutex
		sentTasks         map[string]task
		hc                http.Client
		M                 int
		N                 int
		files             []string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TaskController{
				tasks:             tt.fields.tasks,
				registeredWorkers: tt.fields.registeredWorkers,
				workersQueue:      tt.fields.workersQueue,
				regMutex:          tt.fields.regMutex,
				doneMutex:         tt.fields.doneMutex,
				sentTasks:         tt.fields.sentTasks,
				hc:                tt.fields.hc,
				M:                 tt.fields.M,
				N:                 tt.fields.N,
				files:             tt.fields.files,
			}
			c.createMapTasks()
		})
	}
}

func TestTaskController_createReduceTasks(t *testing.T) {
	type fields struct {
		tasks             chan task
		registeredWorkers map[string]bool
		workersQueue      chan string
		regMutex          *sync.Mutex
		doneMutex         *sync.Mutex
		sentTasks         map[string]task
		hc                http.Client
		M                 int
		N                 int
		files             []string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TaskController{
				tasks:             tt.fields.tasks,
				registeredWorkers: tt.fields.registeredWorkers,
				workersQueue:      tt.fields.workersQueue,
				regMutex:          tt.fields.regMutex,
				doneMutex:         tt.fields.doneMutex,
				sentTasks:         tt.fields.sentTasks,
				hc:                tt.fields.hc,
				M:                 tt.fields.M,
				N:                 tt.fields.N,
				files:             tt.fields.files,
			}
			c.createReduceTasks()
		})
	}
}

func TestTaskController_processTasks(t *testing.T) {
	type fields struct {
		tasks             chan task
		registeredWorkers map[string]bool
		workersQueue      chan string
		regMutex          *sync.Mutex
		doneMutex         *sync.Mutex
		sentTasks         map[string]task
		hc                http.Client
		M                 int
		N                 int
		files             []string
	}
	type args struct {
		isMap bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TaskController{
				tasks:             tt.fields.tasks,
				registeredWorkers: tt.fields.registeredWorkers,
				workersQueue:      tt.fields.workersQueue,
				regMutex:          tt.fields.regMutex,
				doneMutex:         tt.fields.doneMutex,
				sentTasks:         tt.fields.sentTasks,
				hc:                tt.fields.hc,
				M:                 tt.fields.M,
				N:                 tt.fields.N,
				files:             tt.fields.files,
			}
			c.processTasks(tt.args.isMap)
		})
	}
}

func TestTaskController_sendMapTask(t *testing.T) {
	type fields struct {
		tasks             chan task
		registeredWorkers map[string]bool
		workersQueue      chan string
		regMutex          *sync.Mutex
		doneMutex         *sync.Mutex
		sentTasks         map[string]task
		hc                http.Client
		M                 int
		N                 int
		files             []string
	}
	type args struct {
		t   task
		url string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TaskController{
				tasks:             tt.fields.tasks,
				registeredWorkers: tt.fields.registeredWorkers,
				workersQueue:      tt.fields.workersQueue,
				regMutex:          tt.fields.regMutex,
				doneMutex:         tt.fields.doneMutex,
				sentTasks:         tt.fields.sentTasks,
				hc:                tt.fields.hc,
				M:                 tt.fields.M,
				N:                 tt.fields.N,
				files:             tt.fields.files,
			}
			c.sendMapTask(tt.args.t, tt.args.url)
		})
	}
}

func TestTaskController_sendReduceTask(t *testing.T) {
	type fields struct {
		tasks             chan task
		registeredWorkers map[string]bool
		workersQueue      chan string
		regMutex          *sync.Mutex
		doneMutex         *sync.Mutex
		sentTasks         map[string]task
		hc                http.Client
		M                 int
		N                 int
		files             []string
	}
	type args struct {
		t   task
		url string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TaskController{
				tasks:             tt.fields.tasks,
				registeredWorkers: tt.fields.registeredWorkers,
				workersQueue:      tt.fields.workersQueue,
				regMutex:          tt.fields.regMutex,
				doneMutex:         tt.fields.doneMutex,
				sentTasks:         tt.fields.sentTasks,
				hc:                tt.fields.hc,
				M:                 tt.fields.M,
				N:                 tt.fields.N,
				files:             tt.fields.files,
			}
			c.sendReduceTask(tt.args.t, tt.args.url)
		})
	}
}

func Test_getFilenames(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getFilenames(tt.args.dir), "getFilenames(%v)", tt.args.dir)
		})
	}
}
