package group

import "testing"

func BenchmarkGroupIt(b *testing.B){
	var ts []Thing
	for i := 0; i< b.N; i++{
		ts = append(ts, Thing{id: i})
	}
	throttle := 3
	GroupIt(ts, throttle)
}