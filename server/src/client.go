package main

import (
	"fmt"
	"net"
	"time"
	"module"
	"math/rand"
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
				println("Error msgpack unpack:", err.Error())
				continue
			}
        }
	}
}

func client(id int) {
	conn, err := net.Dial("tcp", "127.0.0.1:4313")
	if err != nil {
		fmt.Printf("Error : %s\n", err.Error())
	}
	tc1 := time.Tick(300 * time.Millisecond)
	tc2 := time.Tick(1 * time.Second)
	go ConnService(conn)

	x := id - 20
	y := id - 20
	r := rand.New(rand.NewSource(time.Now().UnixNano() * int64(id)))
	devx := r.Intn(3) - 1
	devy := r.Intn(3) - 1
	loop := 0

	module.SendLoginAsk(conn, "namename" + string(id), "passpass")
	// module.SendRegisterAsk(conn, "grttrtrt", "f9939kk3")
	for {
		select {
		case <- tc1:
			if loop > 10 {
				loop = 0
				devx = r.Intn(3) - 1
				devy = r.Intn(3) - 1
			}
			if x > 120 {
				devx = -1
			} else if x < -120 {
				devx = 1
			}
			if y > 80 {
				devy = -1
			} else if y < -80 {
				devy = 1
			}
			x += devx * 3
			y += devy * 3
			loop++
			module.SendUserMoveAsk(conn, x, y)
		case <- tc2:
			if r.Intn(5) == 0 {
				module.SendPutBombAsk(conn, x, y)
			}
		}
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go client(i)
	}
	client(51)
}
