package module

import(
)

var ChWorld chan *WorldMsg

func goWorld() {
	// state
	var worldState = WorldState{UserCount: 0, MpUserInfo: make(map[int]UserInfoGo)}
	var pState *WorldState = &worldState

	LogDebug("go world running ... ")
	for {
		select {
		case pWorldMsg := <- ChWorld:
			callWorldHandle(pWorldMsg, pState)
		}
	}
}

func WorldStart() {
	ChWorld = make(chan *WorldMsg, 512)
	go goWorld()
}

func callWorldHandle(pWorldMsg *WorldMsg, pWorldState *WorldState) {
	defer func() {
		if x := recover(); x != nil {
			LogDebug("Error from call world handle: ", x)
		}
	}()
	(*pWorldMsg).HandleMsg(pWorldState)
}

// new user in
func (self UserInRpt) HandleMsg(pState *WorldState) {
	pState.UserCount++
	pState.MpUserInfo[self.Userid] = self.UserInfo
	LogDebug("User in, the current user amount: ", pState.UserCount)
}

// new user out
func (self UserOutRpt) HandleMsg(pState *WorldState) {
	pState.UserCount--
	if (pState.UserCount < 0) {
		pState.UserCount = 0
		LogDebug("Why the user amount less then zero?")
	}
	delete(pState.MpUserInfo, self.Userid)
	LogDebug("User out, the current user amount: ", pState.UserCount)
}

// call get user info
func (self UserInfoRpt) HandleMsg(pState *WorldState) {
	if userInfo, ok := pState.MpUserInfo[self.Userid]; ok {  
		self.ChCallBack <- userInfo
	} else {
		self.ChNil <- 1
	}
}