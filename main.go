package main

import (
	"flag"
	"net/http"
	"os"
	"strings"
	"bufio"
)
	

func main() {
	
	var (
		maxQueueSize = flag.Int("max_queue_size", 100, "The size of job queue")
		maxWorkers   = flag.Int("max_workers", 5, "The number of workers to start")
		port         = flag.String("port", "8080", "The server port")
	)
	flag.Parse()
	//tm := NewTemplateManager("login.html","index.html","error.html")//TemplateManager
	var c Connection
	c.Open(readSQLDBConfig("test.properties"))

	jobQueue := make(chan Runner, *maxQueueSize)

	wm := NewWorkerManager(jobQueue, *maxWorkers)
	wm.Run()

	//
	
	// Test work distribution system by pausing each worker for x seconds specified by request.
	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})
	Check(http.ListenAndServe(":"+*port, nil), true)
}

func readSQLDBConfig(filename string) (user string, password string, dbname string, adresse string){
	inFile, err := os.Open(filename)
	Check(err, true)
	defer inFile.Close()
	
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines) 
  
    for scanner.Scan() {
    line := scanner.Text()
    if(strings.Contains(line, "PASSWORD=")){
		    password = strings.Replace(line, "PASSWORD=", "", 1)
    }
    if(strings.Contains(line, "USER=")){
		    user = strings.Replace(line, "USER=", "", 1)
    }
    if(strings.Contains(line, "DATABASE=")){
		    dbname = strings.Replace(line, "DATABASE=", "", 1)
    }
    if(strings.Contains(line, "IP=")){
		    adresse = strings.Replace(line, "IP=", "", 1)
    }
    }
    return
} 