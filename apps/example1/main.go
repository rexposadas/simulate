package main

// This is a sample application which shows a simple use case of
// simulate. This usecase hits the Google homepage once a second -
// that's the default interval behavior.
// The simulate defaults to writing to stdout/stderr.

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rexposadas/simulate"
	simhttp "github.com/rexposadas/simulate/http"
)

type MyActor struct{}

// GetGoogle make a GET request to http://google.com
func (m *MyActor) Run() error {
	resp := simhttp.Get("http://google.com")
	
	return resp
}

func main() {

	simulate.Run()

	// Create job and send to scheduler
	j := simulate.NewJob()
	d := &MyActor{}
	j.Actor = d
	simulate.Jobs <- j
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGQUIT)
	<-sigc
}
