package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/urpc/uio"
)

func main() {
	var port int
	var loops int

	// Example command: go run echo.go --port 9001
	flag.IntVar(&port, "port", 9001, "--port 9001")
	flag.IntVar(&loops, "loops", 0, "num loops")
	flag.Parse()

	var events uio.Events
	events.Pollers = loops
	events.FullDuplex = true
	events.OnStart = func(ev *uio.Events) {
		log.Printf("echo server started on port %d (event-loops: %d)", port, ev.Pollers)
	}
	events.OnData = func(c uio.Conn) error {
		_, err := c.WriteTo(c)
		return err
	}

	log.Fatal(events.Serve(fmt.Sprintf("tcp://0.0.0.0:%d", port)))
}
