package proxy

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/varun-muthanna/forwardproxy/forwardproxypolicy"
)

type handler struct {
	fp *forwardproxypolicy.ForwardProxy
	upstreamAddr string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if h.fp.IsBanned(r.Host) {
		w.Write([]byte("Restricted access to domain " + strconv.Itoa(http.StatusUnauthorized)))
		return
	}

	var url string = "http://" + h.upstreamAddr
	req , err := http.NewRequest(r.Method,url,nil)

	if err!=nil{
		fmt.Printf("Error creating upstream request %s\n",err)
		return 
	}
	
	req.Host=r.Host

	resp ,err := http.DefaultClient.Do(req)

	if err !=nil{
		fmt.Printf("Error sending upstream request %s\n",err)
		return 
	}

	resBody , err := io.ReadAll(resp.Body)

	if err !=nil{
		fmt.Printf("Error in reading upstream reponse %s\n",err)
		return 
	}

	w.Write(resBody)
}

func StartProxy(addr string,fp *forwardproxypolicy.ForwardProxy,upstreamAddr string) {

	h := &handler{
		fp : fp,
		upstreamAddr: upstreamAddr,
	}

	s := http.Server{
		Addr:    addr,
		Handler: h,
	}

	ch := make(chan os.Signal, 1)

	go func() {

		err := s.ListenAndServe()

		if err != nil {
			log.Printf("Server not listening %s \n", err)
		}

	}()

	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)

	sig := <-ch

	fmt.Printf("Initiating gracefull shutdown %s,\n", sig)

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	s.Shutdown(ctx)
}
