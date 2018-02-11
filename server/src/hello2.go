package main

import (
	"bytes"
	"fmt"
	"time"
	"github.com/ugorji/go/codec"
)

func main() {
	fmt.Println("the 1")
	tc:=time.Tick(3 * time.Second)
	for {
    	<-tc
		mh := &codec.MsgpackHandle{RawToString: true}
		data := []interface{}{"abc", 12345, 1.2345, []interface{}{[]interface{}{"ee", 9.6, 33333}, []interface{}{"eges", 1.6, 984}}}
		buf := &bytes.Buffer{}
		enc := codec.NewEncoder(buf, mh)
		enc.Encode(data)

		dec := codec.NewDecoder(buf, mh)
		data1 := make([]interface{}, 10)
		dec.Decode(&data1)
		fmt.Println(data1)
	}
}