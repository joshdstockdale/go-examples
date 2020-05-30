package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Thing struct{
	id int
}
var counter int
var wg sync.WaitGroup
var m sync.Mutex

func main(){
	ts := []Thing{
		{id: 1},{id: 2},{id: 3},{id: 4},{id: 5},{id: 6},{id: 7},{id: 8},{id: 9},{id: 10},
	}

	throttle := 3
	wg.Add(throttle)

	for i:=0; i < throttle; i++{

		//if(runtime.NumGoroutine() > throttle){
		//	wg.Wait()
		//}
		go doit(ts, "throttle")

	}
	wg.Wait()
}
func doit(ts []Thing, s string){
	if s != ""{
		fmt.Println(s)
	}
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	reallydoit(ts)
}
func reallydoit(ts []Thing){
	fmt.Println("Counter\t", counter)
	if counter < len(ts) {
		fmt.Println(ts[counter].id)
	}
	if counter < len(ts) {
		m.Lock()
		counter++
		m.Unlock()
		wg.Add(1)
		doit(ts, "")
	}
	wg.Done()
}