package main

import (
	"bufio"
	"flag"
	"net/http"
	"os"
	"strings"
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

//var c Connection

func init() {
	flag.IntVar(&maxQueueSize, "max_queue_size", 100, "The size of job queue")
	flag.IntVar(&maxWorkers, "max_workers", 5, "The number of workers to start")
	flag.StringVar(&port, "port", "8080", "The server port")
	flag.StringVar(&databaseName, "database_name", "", "The name of the database")
	flag.StringVar(&databaseIP, "database_ip", "localhost", "The IP address of the MySQL server")
	flag.StringVar(&username, "username", "", "The database username")
	flag.StringVar(&password, "password", "", "The database password")
	flag.StringVar(&configFile, "config", "courser.cfg", "The filepath of a config file")
}

func main() {
	flag.Parse()
	//c.Open(username, password, databaseName, databaseIP)

	//Initialize job queue and workers.
	jobQueue := make(chan Runner, maxQueueSize)
	wm := NewWorkerManager(jobQueue, maxWorkers)
	wm.Run()

	//Example HTTP handler for login requests.
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		LoginHandler(w, r, jobQueue)
	})
	Check(http.ListenAndServe(":"+port, nil), true)
}

func readSQLDBConfig(filename string) (user string, password string, dbname string, address string) {
	inFile, err := os.Open(filename)
	Check(err, true)
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "PASSWORD=") {
			password = strings.Replace(line, "PASSWORD=", "", 1)
		}
		if strings.Contains(line, "USER=") {
			user = strings.Replace(line, "USER=", "", 1)
		}
		if strings.Contains(line, "DATABASE=") {
			dbname = strings.Replace(line, "DATABASE=", "", 1)
		}
		if strings.Contains(line, "IP=") {
			address = strings.Replace(line, "IP=", "", 1)
		}
	}
	return
}
