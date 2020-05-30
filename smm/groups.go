package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Thing struct{
	id int
}


var wg sync.WaitGroup
//var m sync.Mutex

func main(){
	ts := []Thing{
		{id: 1},{id: 2},{id: 3},{id: 4},{id: 5},{id: 6},{id: 7},{id: 8},{id: 9},{id: 10},
	}

	throttle := 3
	//wg.Add(throttle)

	var groups [][]Thing
	for counter:=0; counter < len(ts); counter=counter+throttle{
		if counter+throttle > len(ts)-1{
			groups = append(groups, ts[counter:])
			break
		}
		groups = append(groups, ts[counter:counter+throttle])

	}
	fmt.Println("groups", groups)

	for _,g := range groups{
		wg.Add(len(g))

		for _,t := range g{

			go dothing(t)
		}
		wg.Wait()
	}

}
func dothing(t Thing){
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	fmt.Println("Doing the thing...")
	fmt.Println(t.id)
	wg.Done()
}