package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{} // mutual exclusion, a way to evade problems when two processes try to modify data in the same address at the same time
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)        //whenever we spawn a goroutine, we need to add one to the WaitGroup
		go fakeDbCall(i) //now all db calls are run at the same time
	}
	wg.Wait() //wait for the counter to go back to 0 before proceding
	fmt.Printf("\nTotal Execution time: %v", time.Since(t0))
}

func fakeDbCall(i int) {
	//Simulate a db call delay
	var delay int = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock() // this is a full lock. other goroutines won't be able to read or write here while this lock is active
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock() // This prevents other code from reading from results if it's being modified somewhere
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}
