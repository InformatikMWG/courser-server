package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	maxQueueSize int
	maxWorkers   int
	port         string
	databaseName string
	databaseIP   string
	username     string
	password     string
	configFile   string
)

func init() {
	//Setup command line flags.
	flag.IntVar(&maxQueueSize, "max_queue_size", 100, "The size of job queue")
	flag.IntVar(&maxWorkers, "max_workers", 5, "The number of workers to start")
	flag.StringVar(&port, "port", "8080", "The server port")
	flag.StringVar(&databaseName, "database_name", "", "The name of the database")
	flag.StringVar(&databaseIP, "database_ip", "localhost", "The IP address of the MySQL server")
	flag.StringVar(&username, "username", "", "The database username")
	flag.StringVar(&password, "password", "", "The database password")
}

func main() {
	fmt.Println("Done!")
	//Parse command line flags.
	flag.Parse()

	//Initialize job queue and workers.
	Log("Allocating jobQueue channel with buffersize of", maxQueueSize)
	jobQueue := make(chan Runner, maxQueueSize)
	Log("Allocated jobQueue channel, RAM address", jobQueue)
	wm := NewWorkerManager(jobQueue, maxWorkers)
	wm.Run()

	//Example HTTP handler for login requests.
	Log("Registering HTTP handler for /login.")
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		SimpleLoginHandler(w, r, jobQueue)
	})
	Check(http.ListenAndServe(":"+port, nil), true)
}
