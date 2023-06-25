module main

go 1.20

require internal/driver v1.0.0

require github.com/gorilla/mux v1.8.0 // indirect

replace internal/driver => ../../internal/driver
