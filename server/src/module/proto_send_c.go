///////////////////////////////////////////////////////////////////////////
///////// This file is generated by proto/gen_proto.py from ///////////////
///////// proto.txt and proto_data.txt, do not edit!!!		///////////////
///////////////////////////////////////////////////////////////////////////
package module
import (
    "net"
	"fmt"
	"github.com/vmihailenco/msgpack"
)

func SendRegisterAsk(conn net.Conn, Name string, Passwd string) {
	msg := []interface{}{REGISTERASK, Name, Passwd}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendRegisterAsk: ", err.Error())
	} else {
		conn.Write(buf)
	}
}

func SendLoginAsk(conn net.Conn, Name string, Passwd string) {
	msg := []interface{}{LOGINASK, Name, Passwd}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendLoginAsk: ", err.Error())
	} else {
		conn.Write(buf)
	}
}

func SendFriendAsk(conn net.Conn) {
	msg := []interface{}{FRIENDASK}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendFriendAsk: ", err.Error())
	} else {
		conn.Write(buf)
	}
}

func SendRUOKAsk(conn net.Conn, Ok bool) {
	msg := []interface{}{RUOKASK, Ok}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendRUOKAsk: ", err.Error())
	} else {
		conn.Write(buf)
	}
}

func SendUserMoveAsk(conn net.Conn, X int, Y int) {
	msg := []interface{}{USERMOVEASK, X, Y}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendUserMoveAsk: ", err.Error())
	} else {
		conn.Write(buf)
	}
}

func SendAttackAsk(conn net.Conn, Entityid int) {
	msg := []interface{}{ATTACKASK, Entityid}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendAttackAsk: ", err.Error())
	} else {
		conn.Write(buf)
	}
}

func SendPutBombAsk(conn net.Conn, X int, Y int) {
	msg := []interface{}{PUTBOMBASK, X, Y}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendPutBombAsk: ", err.Error())
	} else {
		conn.Write(buf)
	}
}

