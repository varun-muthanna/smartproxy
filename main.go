package main

import (
	"flag"
	"fmt"

	"github.com/varun-muthanna/forwardproxy/config"
	"github.com/varun-muthanna/forwardproxy/forwardproxypolicy"
	"github.com/varun-muthanna/forwardproxy/proxy"
)

func main() {

	configPath := flag.String("config", "config.json", "Path to config file")
	flag.Parse()

	cfg ,err := config.LoadConfig(*configPath)

	if err!=nil{
		fmt.Printf("Error loading config file %s\n",err)
		return 
	}

	fp := forwardproxypolicy.NewForwardProxy(cfg.BannedDomains)

	fmt.Printf("Forward proxy listening on: %s\n", cfg.ListenAddr)
	proxy.StartProxy(cfg.ListenAddr,fp,cfg.UpstreamAddr)
}
