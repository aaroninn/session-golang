# session-golang
自己写的一个简易的session包。  
使用读写锁确保并发安全。  
使用两种方法清除过期的session一种是当访问到过期session时自动清除，另一个是定期扫描map中的session，清除过期session。 
可以刷新session。  
目前设置session过期时间为 24h 扫描时间 1h可以在代码里修改  
使用示例： 
```golang
//初始化
sessions := sessiongo.NewSessionsStorage()
//加入新的session
session := sessiongo.NewSession("test")
session.SetData("test data")
sessions.Add(session)
//备份
sessions.Backup()
//读取备份,备份会覆盖储存之前的sessions，一般用于服务器崩溃后重新启动服务
sessions.ReadBackup()
```



