package pool

import "testing"

func BenchmarkPoolIt(b *testing.B){
	var ts []Thing
	for i := 0; i< b.N; i++{
		ts = append(ts, Thing{id: i})
	}
	throttle := 3
	PoolIt(ts, throttle)
}