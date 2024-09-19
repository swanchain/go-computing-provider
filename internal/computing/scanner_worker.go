package computing

import (
	"sync"
)

func worker(taskQueue <-chan *TaskManagerContract, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskQueue {
		task.Scan()
	}
}

func startWorkerPool(poolSize int) (chan<- *TaskManagerContract, *sync.WaitGroup) {
	taskQueue := make(chan *TaskManagerContract)
	var wg sync.WaitGroup

	for i := 1; i <= poolSize; i++ {
		wg.Add(1)
		go worker(taskQueue, &wg)
	}

	return taskQueue, &wg
}

func submitTask(taskQueue chan<- *TaskManagerContract, task *TaskManagerContract) {
	taskQueue <- task
}
