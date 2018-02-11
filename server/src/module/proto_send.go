///////////////////////////////////////////////////////////////////////////
///////// This file is generated by proto/gen_proto.py from ///////////////
///////// proto.txt and proto_data.txt, do not edit!!!		///////////////
///////////////////////////////////////////////////////////////////////////
package module
import (
	"fmt"
	"github.com/vmihailenco/msgpack"
)

func SendRegisterRsp(chSender chan []byte, OK bool) {
	msg := []interface{}{REGISTERRSP, OK}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendRegisterRsp: ", err.Error())
	} else {
		chSender <- buf
	}
}

func BroRegisterRsp(pState *MapState, OK bool) {
	msg := []interface{}{REGISTERRSP, OK}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendRegisterRsp: ", err.Error())
	} else {
		for _, userInfo := range pState.MpUserInfo {
			userInfo.ChSender <- buf
		}
	}
}

func SendLoginRsp(chSender chan []byte, OK bool, Userid int) {
	msg := []interface{}{LOGINRSP, OK, Userid}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendLoginRsp: ", err.Error())
	} else {
		chSender <- buf
	}
}

func BroLoginRsp(pState *MapState, OK bool, Userid int) {
	msg := []interface{}{LOGINRSP, OK, Userid}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendLoginRsp: ", err.Error())
	} else {
		for _, userInfo := range pState.MpUserInfo {
			userInfo.ChSender <- buf
		}
	}
}

func SendFriendRsp(chSender chan []byte, Ok bool, IdList []interface{}, FriendList []interface{}) {
	msg := []interface{}{FRIENDRSP, Ok, IdList, FriendList}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendFriendRsp: ", err.Error())
	} else {
		chSender <- buf
	}
}

func BroFriendRsp(pState *MapState, Ok bool, IdList []interface{}, FriendList []interface{}) {
	msg := []interface{}{FRIENDRSP, Ok, IdList, FriendList}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendFriendRsp: ", err.Error())
	} else {
		for _, userInfo := range pState.MpUserInfo {
			userInfo.ChSender <- buf
		}
	}
}

func SendRUOKRsp(chSender chan []byte, OK bool) {
	msg := []interface{}{RUOKRSP, OK}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendRUOKRsp: ", err.Error())
	} else {
		chSender <- buf
	}
}

func BroRUOKRsp(pState *MapState, OK bool) {
	msg := []interface{}{RUOKRSP, OK}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendRUOKRsp: ", err.Error())
	} else {
		for _, userInfo := range pState.MpUserInfo {
			userInfo.ChSender <- buf
		}
	}
}

func SendUserIntoMapRsp(chSender chan []byte, Mapid int, X int, Y int, EntityList []interface{}) {
	msg := []interface{}{USERINTOMAPRSP, Mapid, X, Y, EntityList}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendUserIntoMapRsp: ", err.Error())
	} else {
		chSender <- buf
	}
}

func BroUserIntoMapRsp(pState *MapState, Mapid int, X int, Y int, EntityList []interface{}) {
	msg := []interface{}{USERINTOMAPRSP, Mapid, X, Y, EntityList}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendUserIntoMapRsp: ", err.Error())
	} else {
		for _, userInfo := range pState.MpUserInfo {
			userInfo.ChSender <- buf
		}
	}
}

func SendEntityIntoMapRsp(chSender chan []byte, Entity []interface{}) {
	msg := []interface{}{ENTITYINTOMAPRSP, Entity}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendEntityIntoMapRsp: ", err.Error())
	} else {
		chSender <- buf
	}
}

func BroEntityIntoMapRsp(pState *MapState, Entity []interface{}) {
	msg := []interface{}{ENTITYINTOMAPRSP, Entity}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendEntityIntoMapRsp: ", err.Error())
	} else {
		for _, userInfo := range pState.MpUserInfo {
			userInfo.ChSender <- buf
		}
	}
}

func SendEntityOutMapRsp(chSender chan []byte, Entityid int) {
	msg := []interface{}{ENTITYOUTMAPRSP, Entityid}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendEntityOutMapRsp: ", err.Error())
	} else {
		chSender <- buf
	}
}

func BroEntityOutMapRsp(pState *MapState, Entityid int) {
	msg := []interface{}{ENTITYOUTMAPRSP, Entityid}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendEntityOutMapRsp: ", err.Error())
	} else {
		for _, userInfo := range pState.MpUserInfo {
			userInfo.ChSender <- buf
		}
	}
}

func SendEntityPosRsp(chSender chan []byte, EntityList []interface{}) {
	msg := []interface{}{ENTITYPOSRSP, EntityList}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendEntityPosRsp: ", err.Error())
	} else {
		chSender <- buf
	}
}

func BroEntityPosRsp(pState *MapState, EntityList []interface{}) {
	msg := []interface{}{ENTITYPOSRSP, EntityList}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendEntityPosRsp: ", err.Error())
	} else {
		for _, userInfo := range pState.MpUserInfo {
			userInfo.ChSender <- buf
		}
	}
}

func SendEntityHpRsp(chSender chan []byte, Entityid int, Hp int) {
	msg := []interface{}{ENTITYHPRSP, Entityid, Hp}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendEntityHpRsp: ", err.Error())
	} else {
		chSender <- buf
	}
}

func BroEntityHpRsp(pState *MapState, Entityid int, Hp int) {
	msg := []interface{}{ENTITYHPRSP, Entityid, Hp}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendEntityHpRsp: ", err.Error())
	} else {
		for _, userInfo := range pState.MpUserInfo {
			userInfo.ChSender <- buf
		}
	}
}

func SendBombInRsp(chSender chan []byte, Bombid int, X int, Y int) {
	msg := []interface{}{BOMBINRSP, Bombid, X, Y}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendBombInRsp: ", err.Error())
	} else {
		chSender <- buf
	}
}

func BroBombInRsp(pState *MapState, Bombid int, X int, Y int) {
	msg := []interface{}{BOMBINRSP, Bombid, X, Y}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendBombInRsp: ", err.Error())
	} else {
		for _, userInfo := range pState.MpUserInfo {
			userInfo.ChSender <- buf
		}
	}
}

func SendBombBlastRsp(chSender chan []byte, Bombid int) {
	msg := []interface{}{BOMBBLASTRSP, Bombid}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendBombBlastRsp: ", err.Error())
	} else {
		chSender <- buf
	}
}

func BroBombBlastRsp(pState *MapState, Bombid int) {
	msg := []interface{}{BOMBBLASTRSP, Bombid}
	buf, err := msgpack.Marshal(msg)
	if err != nil {
		fmt.Println("Error in SendBombBlastRsp: ", err.Error())
	} else {
		for _, userInfo := range pState.MpUserInfo {
			userInfo.ChSender <- buf
		}
	}
}

