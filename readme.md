## 1. 安装依赖包
```swift

go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get github.com/go-sql-driver/mysql
go get golang.org/x/crypto/bcrypt
```

## 2. 实现功能

+ 注册：判断用户名是否存在 不存在则hash密码 存入db 注册成功
+ 登录 ：判断用户名是否存在 存在则判断密码是否正确 成功登录
+ 删除：判断用户名是否存在 存在则删除
+ 更新密码：判断用户名是否存在 存在则更新密码

## 3. 代码架构
+ api
    + database--mysql.go gorm连接到本地mysql
    + models--user.go 创建User结构体及对应的数据库操作函数
    + apis--user.go 处理请求并响应
    + router--router.go 路由
+ main.go 启动server 默认8080