package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Simple sample app in Go for a basic API server
func main() {

	// Handle GET methods for fetching the current Unix EPOCH
	http.HandleFunc("/time", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Simple struct to hold our API response
		type timeAPI struct {
			Payload int64
			Status  string
		}

		// Build the payload
		data := timeAPI{Payload: time.Now().Unix(), Status: "OK"}

		// Output the JSON to the client
		json.NewEncoder(w).Encode(data)

	})

	// Handle GET/POST requests for storing a log request
	http.HandleFunc("/log", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Simple struct to hold our API response
		type logAPI struct {
			Status string
		}

		// Fetch the logdata (via a GET or POST request)
		logdata := req.FormValue("logdata")

		// Append the output to our sample iot-sensor.log, create the file if it does not exist.
		f, err := os.OpenFile("/tmp/iot-sensor.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		// Write the `logdata` from our request to the file with the current EPOCH
		if _, err := f.Write([]byte(fmt.Sprintf("%d => %s\n", time.Now().Unix(), logdata))); err != nil {
			log.Fatal(err)
		}

		// Close the file
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}

		// Output to the client the `OK` JSON response
		data := logAPI{Status: "OK"}
		json.NewEncoder(w).Encode(data)

	})

	// Serve our sample web-server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	/*

			How go uses concurrency to launch multiple threads, from server.go

		   // Serve accepts incoming connections on the Listener l, creating a
		   // new service goroutine for each.  The service goroutines read requests and
		   // then call srv.Handler to reply to them.
		   func (srv *Server) Serve(l net.Listener) error {
		           for {
		               rw, e := l.Accept()
		                  if e != nil {
		   ......
		               c, err := srv.newConn(rw)
		               if err != nil {
		                   continue
		               }
		               c.setState(c.rwc, StateNew) // before Serve can return
		               go c.serve() // Here, baby!
		           }
		   }

	*/

}
