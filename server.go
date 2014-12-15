package simulate

import (
	"fmt"
)

var (
	Jobs  chan *Job
	Port  int
	Stats *StatsObj
)

// Run runs the simulate server
func Run() {
	Jobs = make(chan *Job, 1000)
	Stats = &StatsObj{}
	go Stats.Run()

	fmt.Println("Simulator started \n\n")

	go worker()
}

// Worker runs the jobs.
// Each job received is a new routine.
func worker() {

	for {
		j := <-Jobs
		go j.Actor.Run()
	}
}
