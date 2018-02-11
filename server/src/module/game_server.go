package module

import (
	"net"
)


// the go for sending msg
func goSender(conn net.Conn, chSender chan []byte, chClose chan byte) {
	for {
		select {
		case buf := <-chSender:
			_, err := conn.Write(buf)
			if err != nil {
				LogDebug("Error writing:", err.Error())
			}
		case <- chClose:
			close(chClose)
			return
		}
	}
}

// the go for reading msg and router msg
func goReader(conn net.Conn, chSender chan []byte, chClose chan byte) {
	// state
	var connState = ConnState{ChSender: chSender}
	var pConnState *ConnState = &connState
	buf := make([]byte, 1024)

	// msg decode data
	var data []interface{}

	// defer
	defer func() {
		conn.Close()
		close(chSender)
		chClose <- 1
		if pConnState.UserInfo.ChUser != nil {
			pConnState.UserInfo.ChUser <- &ConnClosed{}
		}
	}()

	for {
		len, err := conn.Read(buf)
		if err != nil {
			LogDebug("Error reading:", err.Error())
			return
		}
		if len > 0 {
			Recv(pConnState, buf, data)
		}
	}
}

func GameStart() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:4313")
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	LogDebug("server running ...")

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			LogDebug("Error accept:", err.Error())
			continue
		}
		LogDebug("Accept from:", conn.RemoteAddr().String())
		chSender := make(chan []byte, 1024)
		chClose := make(chan byte, 1)
		go goReader(conn, chSender, chClose)
		go goSender(conn, chSender, chClose)
     }
}
