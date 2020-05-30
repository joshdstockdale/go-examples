package pool

import (
	"fmt"
	"sync"
)

type Thing struct{
	id int
}
var counter int
var wg sync.WaitGroup
var m sync.Mutex

func PoolIt(ts []Thing, throttle int){

	wg.Add(throttle)

	for i:=0; i < throttle; i++{
		go doit(ts, "throttle")
	}
	wg.Wait()
}
func doit(ts []Thing, s string){
	reallydoit(ts)
}
func reallydoit(ts []Thing){
	if counter < len(ts) {
		fmt.Printf("%v,",ts[counter].id)
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