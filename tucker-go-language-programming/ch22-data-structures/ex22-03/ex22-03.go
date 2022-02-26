package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	data *list.List
}

func (stack *Stack) Push(value interface{}) {
	stack.data.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	back := stack.data.Back()
	if back != nil {
		return stack.data.Remove(back)
	}
	return nil
}

func InitStack() *Stack {
	return &Stack{list.New()}
}

func main() {
	stack := InitStack()

	for i := 1; i < 5; i++ {
		stack.Push(i)
	}

	value := stack.Pop()
	for value != nil {
		fmt.Printf("%v -> ", value)
		value = stack.Pop()
	}
}
