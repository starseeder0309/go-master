package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	data *list.List
}

func (queue *Queue) Push(value interface{}) {
	queue.data.PushBack(value)
}

func (queue *Queue) Pop() interface{} {
	front := queue.data.Front()
	if front != nil {
		return queue.data.Remove(front)
	}
	return nil
}

func InitQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	queue := InitQueue()

	for i := 1; i < 5; i++ {
		queue.Push(i)
	}

	value := queue.Pop()
	for value != nil {
		fmt.Printf("%v -> ", value)
		value = queue.Pop()
	}
}
