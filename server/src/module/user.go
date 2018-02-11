package module

import (
	"time"
)

// the go for dealing user stuff
func goUser(pConnState *ConnState, user User) {
	// channel
	chReConn := make(chan ReconnInfo, 1)
	chUser := make(chan UserMsg, 512)
	chUserExit := make(chan byte, 1)

	defer func() {
		close(chUser)
		close(chReConn)
		close(chUserExit)
	}()

	// user state
	var userState = UserState{User: user, ChUserExit: chUserExit, ChSender: pConnState.ChSender, ChReConn: chReConn}
	var pUserState *UserState = &userState

	defer handleUserExit(pUserState)

	// conn state
	var info = UserInfoGo{ChUser: chUser, ChUserExit: chUserExit, ChReConn: chReConn}
	pConnState.UserInfo = info

	// tell world
	var worldMsg WorldMsg = UserInRpt{Userid: pUserState.User.Userid, UserInfo: info}
	ChWorld <- &worldMsg

	// tell map
	UserInMap(&pUserState.User, pUserState.ChSender)

	LogDebug("go user running ... ")
	for {
		select {
		case userMsg := <-chUser:
			callUserHandle(userMsg, pUserState)
		case reconnInfo := <- chReConn:
			// 顶号
			pUserState.ChSender = reconnInfo.ChSender
			SendLoginRsp(pUserState.ChSender, true, pUserState.User.Userid)
			LogDebug("user displace ok")
		case <- chUserExit:
			LogDebug("go user exit")
			return
		}
	}
}

func callUserHandle(msg UserMsg, pUserState *UserState) {
	defer func() {
		if x := recover(); x != nil {
			LogDebug("Error from call user handle: ", x)
		}
	}()
	msg.HandleMsg(pUserState)
	msg.Release()
}

// register
func UserRegister(pConnState *ConnState, ask RegisterAsk) {
	name := ask.Name
	passwd := ask.Passwd

	has, err := DBHasUserName(name)
	if err != nil {
		SendRegisterRsp(pConnState.ChSender, false)
		return
	}

	if has {
		SendRegisterRsp(pConnState.ChSender, false)
		return
	}

	user, err := DBNewUser(name, passwd)
	if err != nil {
		SendRegisterRsp(pConnState.ChSender, false)
		return
	}

	LogDebug("user register success: ", user)
	SendRegisterRsp(pConnState.ChSender, true)
	LoadUser(pConnState, user)
}

// login
func UserLogin(pConnState *ConnState, ask LoginAsk) {
	// 此处检查账号密码
	name := ask.Name
	passwd := ask.Passwd

	has, err := DBHasUserName(name)
	if (err != nil) {
		SendLoginRsp(pConnState.ChSender, false, 0)
		return
	}
	if (!has) {
		SendLoginRsp(pConnState.ChSender, false, 0)
		return
	}

	user, err := FindUserByName(name)
	if (err != nil) {
		SendLoginRsp(pConnState.ChSender, false, 0)
		return
	}

	if (user.Passwd != passwd) {
		SendLoginRsp(pConnState.ChSender, false, 0)
		return
	}

	LoadUser(pConnState, user)
}

func LoadUser(pConnState *ConnState, user User) {
	// call world to check user go is running
	chCallBack := make(chan UserInfoGo)
	chNil := make(chan byte)

	defer func() {
		close(chCallBack)
		close(chNil)
	}()

	var rpt WorldMsg = UserInfoRpt{Userid: user.Userid, ChCallBack: chCallBack, ChNil: chNil}
	ChWorld <- &rpt
	select {
		// 断线重连 或 顶号
		case userInfo := <- chCallBack:
			LogDebug("get userinfo from world: ", userInfo)
			pConnState.UserInfo = userInfo
			pConnState.UserInfo.ChReConn <- ReconnInfo{pConnState.ChSender}
			UserInMap(&user, pConnState.ChSender)
		// 新开go
		case <- chNil:
			go goUser(pConnState, user)
			SendLoginRsp(pConnState.ChSender, true, user.Userid)
		// timeout
		case <- time.After(time.Second * 10):
			LogDebug("get userinfo from world timeout")
			SendLoginRsp(pConnState.ChSender, false, 0)
	}
}

func UserInMap(pUser *User, chSender chan []byte) {
	if pUser.Mapid == 0 {
		pUser.Mapid = 1
		pUser.X = 0
		pUser.Y = 0
	}
	if chMap, ok := MpMapCh[pUser.Mapid]; ok {
		chMap <- &UserInAct{pUser, chSender}
	}
}

func UserOutMap(pUser *User) {
	if chMap, ok := MpMapCh[pUser.Mapid]; ok {
		chMap <- &UserOutAct{pUser.Userid}
	}
}

// conn closed
func (self *ConnClosed) HandleMsg(pUserState *UserState) {
	SaveUser(pUserState.User)
	UserOutMap(&pUserState.User)
	// 等待重连
	select {
		case reconnInfo := <- pUserState.ChReConn:
			// 重连成功
			pUserState.ChSender = reconnInfo.ChSender
			SendLoginRsp(pUserState.ChSender, true, pUserState.User.Userid)
			LogDebug("user reconnect ok")
		// 断线重连等待时间
		case <- time.After(time.Second * 100):
			pUserState.ChUserExit <- 1
	}
}

func (self *ConnClosed) Release() {
	// do nothing
}

// user exit
func handleUserExit(pUserState *UserState) {
	// tell world
	var msg WorldMsg = UserOutRpt{Userid: pUserState.User.Userid}
	ChWorld <- &msg
}

// friend ask
func (self *FriendAsk) HandleMsg(pUserState *UserState) {
	
}

// are you ok
func (self *RUOKAsk) HandleMsg(pUserState *UserState) {
	LogDebug("I am ", self.Ok)
	SendRUOKRsp(pUserState.ChSender, !self.Ok)
}
