package main

import (
	"fmt"
	"net"
	"time"
	"module"
	"github.com/vmihailenco/msgpack"
)

func ConnService(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		len, err := conn.Read(buf)
		if len > 0 {
			var out = make([]interface{}, 1)
			err = msgpack.Unmarshal(buf, &out)
			if err != nil {
				fmt.Println("Error msgpack unpack: ", err.Error())
				continue
			}
			fmt.Println("receive: ", out)
        }
	}
}

func client(id int) {
	conn, err := net.Dial("tcp", "127.0.0.1:4313")
	if err != nil {
		fmt.Printf("Error : %s\n", err.Error())
	}
	tc := time.Tick(100 * time.Millisecond)
	go ConnService(conn)

	a := true
	module.SendLoginAsk(conn, "namename" + string(id), "passpass")
	// module.SendRegisterAsk(conn, "grttrtrt", "f9939kk3")
	for {
		select {
		case <- tc:
			a = !a
			module.SendRUOKAsk(conn, a)
		}
	}
}

func main() {
	client(51)
}
