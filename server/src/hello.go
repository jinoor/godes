package main

import (
	"fmt"
	"time"
	"msgpack"
)

func main() {
	fmt.Println("the 1")
	tc:=time.Tick(3 * time.Second)
	for {
    	<-tc
		in := map[string]interface{}{"foo": 1, "hello": "world", "arr": map[string]interface{}{"a":1, "b":2, "c":3}}
		// in := []interface{}{"abc", 12345, 1.2345, []interface{}{[]interface{}{"ee", 9.6, 33333}, []interface{}{"eges", 1.6, 984}}}
        b, err := msgpack.Marshal(in)
        if err != nil {
            panic(err)
        }

		// var out = make([]interface{}, 1)
		var out map[string]interface{}
        err = msgpack.Unmarshal(b, &out)
        if err != nil {
            panic(err)
        }

		fmt.Println(out)
        // fmt.Println("foo =", out["foo"])
        // fmt.Println("hello =", out["hello"])
		// fmt.Println("arr = ", out["arr"])
	}
}