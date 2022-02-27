package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func dingingProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다.\n", name)
		first.Lock()
		fmt.Printf("%s %s 획득\n", name, firstName)
		second.Lock()
		fmt.Printf("%s %s 획득\n", name, secondName)

		fmt.Printf("%s 밥을 먹습니다.\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}

	waitGroup.Done()
}

func main() {
	rand.Seed(time.Now().UnixMilli())

	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	waitGroup.Add(2)

	go dingingProblem("A", fork, spoon, "포크", "스푼")
	go dingingProblem("B", spoon, fork, "스푼", "포크")
	// go dingingProblem("B", fork, spoon, "포크", "스푼")

	waitGroup.Wait()
}
