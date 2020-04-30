# Mongo操作示例接口

1. 插入一条默认数据

   curl -X GET localhost:7460/operation/insert

2. 插入N条默认数据

   curl -X GET localhost:7460/operation/insert/3
   
3. 查找数据

   + 不限返回数量

     curl -X POST localhost:7460/operation/find -d '{"name":"op"}'

   + 限制返回数量

     curl 









