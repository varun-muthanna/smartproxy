# smartproxy
A proxy (forward + reverse) in Go 

## Features
Forward proxy with host header based forwarding/blocking

## Usage

go run main.go   Make the forward proxy up running

Run the reverse proxy as explained in the repo 

curl http://localhost:8080 -H "host:example.com"   (Hit the forward proxy with Public DNS)

wrk -t4 -c5 -d2s -s host.lua http://localhost:8080  (Hit the forward proxy with local backend servers)
