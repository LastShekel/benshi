module main

go 1.20

require internal/driver v1.0.0

require (
	github.com/gorilla/mux v1.8.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace internal/driver => ../../internal/driver
