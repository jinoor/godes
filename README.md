# godes
简单的游戏框架，后端GO，前端Unity

### 依赖库：msgpack mongodb ###
* go get -u github.com/vmihailenco/msgpack
* go get gopkg.in/mgo.v2

### MongoDB ###
* 安装mongodb，把bin路径添加到环境变量
* 在随便一个盘新建目录如E:\mongodb\db
* 把mgo_start.bta里的路径改成刚刚的路径，然后启动

### 启动 ###
go run server.go

### 测试 ###
go run client.go
