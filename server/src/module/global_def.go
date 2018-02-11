package module

// 玩家的conn goroutine中运行的state
type ConnState struct {
	ChSender chan []byte	// 给玩家socket发送消息的管道
	UserInfo UserInfoGo
}

// 发送给玩家user goroutine的消息抽象接口,客户端发来的消息struct定义在proto_def,需要实现HandleUser方法
type UserMsg interface {
	HandleMsg(pState *UserState)
	Release()
}

// 玩家的user goroutine中运行的state
type UserState struct {
	User User
	ChSender chan []byte		// 给玩家socket发送消息的管道
	ChReConn chan ReconnInfo	// 等待重连channel
	ChUserExit chan byte 		// 通知user go退出
}

// 玩家连接断开通知user go
type ConnClosed struct {

}

// 玩家重连信息
type ReconnInfo struct {
	ChSender chan []byte
}

// 世界world goroutine中运行的state
type WorldState struct {
	UserCount int						// 玩家数
	MpUserInfo map[int]UserInfoGo	// 玩家id到玩家信息的索引
}

// 玩家的一些chennel信息
type UserInfoGo struct {
	ChUser chan UserMsg		// 给玩家socket发送消息的管道
	ChUserExit chan byte 		// 通知user go退出
	ChReConn chan ReconnInfo	// 等待重连channel
}

// 发送给world goroutine的消息抽象接口
type WorldMsg interface {
	HandleMsg(pState *WorldState)
}

// 通知world新的用户进来
type UserInRpt struct {
	Userid int
	UserInfo UserInfoGo
}

// 通知world用户离开
type UserOutRpt struct {
	Userid int
}

// 向world请求玩家channle
type UserInfoRpt struct {
	Userid int
	ChCallBack chan UserInfoGo
	ChNil chan byte
}

// 地图map goroutine中运行的state
type MapState struct {
	Mapid int						// 地图id
	MpUserInfo map[int]UserInfoMap	// 玩家id到玩家信息的索引
	MpMon map[int]*MapMon			// 地图中的怪物
	MpBomb map[int]*MapBomb			// 地图中的炸弹
}

// 地图中的怪物
type MapMon struct {
	Monid int
	X, Y int
	Hp int
	DevX, DevY int
	LoopCount int
}

// 地图中的炸弹
type MapBomb struct {
	Bombid int
	Ownerid int
	X, Y int
	CreateTime int64
}

// 发送给地图map goroutine的消息抽象接口
type MapMsg interface {
	HandleMsg(pState *MapState)
	Release()
}

// 玩家进入地图
type UserInAct struct {
	User *User
	ChSender chan []byte
}

// 玩家离开地图
type UserOutAct struct {
	Userid int
}

// 玩家地图信息
type UserInfoMap struct {
	User *User
	ChSender chan []byte
}
