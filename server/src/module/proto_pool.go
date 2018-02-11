///////////////////////////////////////////////////////////////////////////
///////// This file is generated by proto/gen_proto.py from ///////////////
///////// proto.txt and proto_data.txt, do not edit!!!		///////////////
///////////////////////////////////////////////////////////////////////////
package module

import(
	"sync"
)

var RegisterAskPool *sync.Pool
var LoginAskPool *sync.Pool
var FriendAskPool *sync.Pool
var RUOKAskPool *sync.Pool
var UserMoveAskPool *sync.Pool
var AttackAskPool *sync.Pool
var PutBombAskPool *sync.Pool
var MapUserMoveAskPool *sync.Pool
var MapAttackAskPool *sync.Pool
var MapPutBombAskPool *sync.Pool

func InitProtoPool() {
	RegisterAskPool = &sync.Pool{
		New: func() interface{} {
			return RegisterAsk{}
		},
	}

	LoginAskPool = &sync.Pool{
		New: func() interface{} {
			return LoginAsk{}
		},
	}

	FriendAskPool = &sync.Pool{
		New: func() interface{} {
			return FriendAsk{}
		},
	}

	RUOKAskPool = &sync.Pool{
		New: func() interface{} {
			return RUOKAsk{}
		},
	}

	UserMoveAskPool = &sync.Pool{
		New: func() interface{} {
			return UserMoveAsk{}
		},
	}

	AttackAskPool = &sync.Pool{
		New: func() interface{} {
			return AttackAsk{}
		},
	}

	PutBombAskPool = &sync.Pool{
		New: func() interface{} {
			return PutBombAsk{}
		},
	}

	MapUserMoveAskPool = &sync.Pool{
		New: func() interface{} {
			return MapUserMoveAsk{}
		},
	}

	MapAttackAskPool = &sync.Pool{
		New: func() interface{} {
			return MapAttackAsk{}
		},
	}

	MapPutBombAskPool = &sync.Pool{
		New: func() interface{} {
			return MapPutBombAsk{}
		},
	}

}

func (self *RegisterAsk) Release() {
	RegisterAskPool.Put(*self)
}

func (self *LoginAsk) Release() {
	LoginAskPool.Put(*self)
}

func (self *FriendAsk) Release() {
	FriendAskPool.Put(*self)
}

func (self *RUOKAsk) Release() {
	RUOKAskPool.Put(*self)
}

func (self *UserMoveAsk) Release() {
	UserMoveAskPool.Put(*self)
}

func (self *AttackAsk) Release() {
	AttackAskPool.Put(*self)
}

func (self *PutBombAsk) Release() {
	PutBombAskPool.Put(*self)
}

func (self *MapUserMoveAsk) Release() {
	MapUserMoveAskPool.Put(*self)
}

func (self *MapAttackAsk) Release() {
	MapAttackAskPool.Put(*self)
}

func (self *MapPutBombAsk) Release() {
	MapPutBombAskPool.Put(*self)
}
