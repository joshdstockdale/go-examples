package group

import (
	"fmt"
	"sync"
)

type Thing struct{
	id int
}

var wg sync.WaitGroup
//var m sync.Mutex

func GroupIt(ts []Thing, throttle int){

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

	fmt.Printf("%v,",t.id)
	wg.Done()
}