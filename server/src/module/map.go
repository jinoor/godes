package module

import(
	"time"
	"math/rand"
)

const USER = 1
const MONSTER = 2

var MpMapCh map[int]chan MapMsg
var MonCount int = 1000001
var BombCount int = 2000001

func goMap(mapid int) {
	// state
	var mapState = MapState{Mapid: mapid, MpUserInfo: make(map[int]UserInfoMap), MpMon: make(map[int]*MapMon), MpBomb: make(map[int]*MapBomb)}
	var pState *MapState = &mapState

	// map channel
	var ChMap = make(chan MapMsg, 512)
	MpMapCh[mapid] = ChMap

	// monster
	initMon(pState)

	// loop tick
	tcLoop := time.Tick(200 * time.Millisecond)

	LogDebug("go map ", mapid, " running ... ")
	for {
		select {
		case mapMsg := <- ChMap:
			callMapHandle(mapMsg, pState)
		case <- tcLoop:
			loop(pState)
		}
	}
}

func MapStart() {
	MpMapCh = make(map[int]chan MapMsg, 10)
	for i := 0; i < 10; i++ {
		go goMap(i)
	}
}

func callMapHandle(msg MapMsg, pMapState *MapState) {
	defer func() {
		if x := recover(); x != nil {
			LogDebug("Error from call map handle: ", x)
		}
	}()
	msg.HandleMsg(pMapState)
	msg.Release()
}

func initMon(pMapState *MapState) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 2; i++ {
		x := r.Intn(240) - 120
		y := r.Intn(160) - 80
		devx := r.Intn(3) - 1
		devy := r.Intn(3) - 1
		var mon = MapMon{Monid: MonCount, X: x, Y: y, DevX: devx, DevY:devy, Hp:10, LoopCount: 0}
		pMapState.MpMon[MonCount] = &mon
		MonCount++
	}
}

// map loop
func loop(pState *MapState) {
	defer func() {
		if x := recover(); x != nil {
			LogDebug("Error from map loop: ", x)
		}
	}()
	lsMon := make([]interface{}, len(pState.MpMon))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	i := 0
	for monid, pMon := range pState.MpMon {
		if pMon.LoopCount > 10 {
			pMon.DevX += r.Intn(3) - 1
			pMon.DevY += r.Intn(3) - 1
			if pMon.X > 120 {
				pMon.DevX = -1
			} else if pMon.X < -120 {
				pMon.DevX = 1
			}
			if pMon.Y > 80 {
				pMon.DevY = -1
			} else if pMon.Y < -80 {
				pMon.DevY = 1
			}
		}
		pMon.X += pMon.DevX
		pMon.Y += pMon.DevY
		pMon.LoopCount++
		lsMon[i] = []interface{}{monid, MONSTER, pMon.X, pMon.Y}
		i++
	}

	// brocast
	BroEntityPosRsp(pState, lsMon)

	now := time.Now().Unix()
	for bombid, pBomb := range pState.MpBomb {
		if now - pBomb.CreateTime >= 2 {
			BroBombBlastRsp(pState, bombid)
			delete(pState.MpBomb, bombid)
		}
	}
}

// 玩家进入地图
func (self *UserInAct) HandleMsg(pState *MapState) {
	pState.MpUserInfo[self.User.Userid] = UserInfoMap{self.User, self.ChSender}
	lsEntity := make([]interface{}, len(pState.MpUserInfo) - 1 + len(pState.MpMon))
	i := 0
	// other user
	for otherid, otherInfo := range pState.MpUserInfo {
		if self.User.Userid != otherid {
			lsEntity[i] = []interface{}{otherid, USER, otherInfo.User.X, otherInfo.User.Y}
			SendEntityIntoMapRsp(otherInfo.ChSender, []interface{}{[]interface{}{self.User.Userid, USER, self.User.X, self.User.Y}})
			i++
		}
	}
	// monster
	for monid, pMon := range pState.MpMon {
		lsEntity[i] = []interface{}{monid, MONSTER, pMon.X, pMon.Y}
		i++
	}
	SendUserIntoMapRsp(self.ChSender, self.User.Mapid, self.User.X, self.User.Y, lsEntity)

	LogDebug("User into map ", pState.Mapid, ", curr map user amount: ", len(pState.MpUserInfo))
}

func (self *UserInAct) Release() {
	
}

func (self *UserOutAct) HandleMsg(pState *MapState) {
	delete(pState.MpUserInfo, self.Userid)
	BroEntityOutMapRsp(pState, self.Userid)
	LogDebug("User out map ", pState.Mapid, ", curr map user amount: ", len(pState.MpUserInfo))
}

func (self *UserOutAct) Release() {
	
}

func (self *MapUserMoveAsk) HandleMsg(pState *MapState) {
	if userInfo, ok := pState.MpUserInfo[self.Userid]; ok {
		userInfo.User.X = self.X
		userInfo.User.Y = self.Y
		// 广播玩家
		BroEntityPosRsp(pState, []interface{}{[]interface{}{self.Userid, USER, self.X, self.Y}})
	}
}

// user attack
func (self MapAttackAsk) HandleMsg(pState *MapState) {
	if pMon, ok := pState.MpMon[self.Entityid]; ok {
		pMon.Hp -= 2
		if pMon.Hp > 0 {
			BroEntityHpRsp(pState, pMon.Monid, pMon.Hp)
		} else {
			delete(pState.MpMon, pMon.Monid)
			BroEntityOutMapRsp(pState, pMon.Monid)
		}
	}
}

// user put bomb
func (self MapPutBombAsk) HandleMsg(pState *MapState) {
	if _, ok := pState.MpUserInfo[self.Userid]; ok {
		bombid := BombCount
		BombCount++
		var bomb = MapBomb{bombid, self.Userid, self.X, self.Y, time.Now().Unix()}
		pState.MpBomb[bombid] = &bomb
		BroBombInRsp(pState, bomb.Bombid, bomb.X, bomb.Y)
	}
}


