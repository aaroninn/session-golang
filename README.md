# session-golang
自己写的一个简易的session包。  
使用读写锁确保并发安全。  
使用两种方法清除过期的session一种是当访问到过期session时自动清除，另一个是定期扫描map中的session，清除过期session。 
可以设置当session被访问时自动续期。  
目前设置session过期时间为 24h 扫描时间 1h可以在代码里修改  
