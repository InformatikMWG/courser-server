package main

// Runner is the interface that wraps the run() method.
type Runner interface {
	run() error
}

// newWorker creates a new worker.
func newWorker(jobQueue chan Runner) {
	go func() {
		for {
			select {
			case job := <-jobQueue:
				job.run()
			}
		}
	}()
}

// NewWorkerManager creates a new work dispatcher.
func NewWorkerManager(jobQueue chan Runner, maxWorkers int) *WorkerManager {
	Log("Initializing worker manager for queue", jobQueue, "with", maxWorkers, "parallel workers.")
	return &WorkerManager{maxWorkers, jobQueue}
}

// WorkerManager distributes work and manages parallel workers.
type WorkerManager struct {
	maxWorkers int
	jobQueue   chan Runner
}

// Run initializes the worker manager and starts it.
func (d *WorkerManager) Run() {
	// Create and start workers.
	Log("Creating workers for channel", d.jobQueue)
	for i := 0; i < d.maxWorkers; i++ {
		newWorker(d.jobQueue)
		Log("Created worker", i+1, "/", d.maxWorkers)
	}
}
