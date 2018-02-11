package module

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session
var collect *mgo.Collection

func DBStart() {
	s, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	session = s

	session.SetMode(mgo.Monotonic, true)

	collect = session.DB("godes").C("user")
	LogDebug("db session running ...")
}

func DBNewUser(name, passwd string) (user User, err error) {
	count, err := collect.Count()
	if err != nil {
		LogDebug("db new user error: ", err)
		return user, err
	}

	userid := 0
	if count == 0 {
		userid = 100001
	} else {
		userid = count + 100001
	}

	user = User{userid, name, passwd, 1, 0, 0}
	err = collect.Insert(&user)
	if err != nil {
		LogDebug("db new user error: ", err)
		return user, err
	}

	return user, nil
}

// 通过name查找User
func FindUserByName(name string) (user User, err error) {
	err = collect.Find(bson.M{"name": name}).One(&user)
	if err != nil {
		LogDebug("find user by name error: ", err)
		return user, err
	}

	return user, nil
}

// 通过id查找User
func FindUserById(userid int) (user User, err error) {
	err = collect.Find(bson.M{"userid": userid}).One(&user)
	if err != nil {
		LogDebug("find user by userid error: ", err)
		return user, err
	}

	return user, nil
}

func DBHasUserName(name string) (bool, error) {
	count, err := collect.Find(bson.M{"name": name}).Count()
	if err != nil {
		LogDebug("check has user name error: ", err)
		return false, err
	}

	if (count > 0) {
		return true, nil
	} else {
		return false, nil
	}
}

// save
func SaveUser(user User) (bool, error) {
	err := collect.Update(bson.M{"userid": user.Userid}, user)
	if err != nil {
		LogDebug("save user error: ", err)
		return false, err
	}

	LogDebug("save user success: ", user)
	return true, nil
}