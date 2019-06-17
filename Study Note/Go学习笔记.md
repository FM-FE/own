## Go学习笔记

1. “ ... ” 意味着，可变长参数

   e.g. 

   声明 ：`func (s *Server) InitRouter(routers ...router.Router)`

   调用：`server.InitRouter(routers...)`

2. 如何使用 var 定义的变量