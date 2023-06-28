# benshi
Distributed word count for Benshi

To build driver use

``go build -C ./cmd/driver -o driver``

To build worker 

``go build -C ./cmd/worker -o worker``

How to run

``driver -M 8 -N 8``

this will start driver it will look up to ./configs/main.yml for settings

To start wokrer you just should run

``worker``

you have 20 seconds to start as many workers, as you want.

Workers will start to register in driver.
