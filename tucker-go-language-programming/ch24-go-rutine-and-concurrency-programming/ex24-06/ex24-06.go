package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (job *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", job.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", job.index, job.index*job.index)
}

func main() {
	var jobs [10]Job

	for i := 0; i < 10; i++ {
		jobs[i] = &SquareJob{i}
	}

	waitGroup.Add(10)

	for i := 0; i < 10; i++ {
		job := jobs[i]
		go func() {
			job.Do()
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}
