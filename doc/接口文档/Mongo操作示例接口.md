# Mongo操作示例接口

1. 插入默认数据

   + 插入一条

     curl -X GET localhost:7460/operation/insert

   + 插入N条

     curl -X GET localhost:7460/operation/insert/3

2. 查找数据

   + 查找一条

     curl -X POST localhost:7460/operation/find/one -d '{"key":"name","value":"op"}' 
     
   + 查找所有

     curl -X POST localhost:7460/operation/find -d '{"key":"name","value":"op"}'

3. 更新数据

   + 更新一条

     curl -X POST localhost:7460/operation/update/one -d '{"filter":{"key":"name","value":"op"},"update":{"key":"description","value":"update"}}'

   + 更新所有

     curl -X POST localhost:7460/operation/update -d '{"filter":{"key":"name","value":"op"},"update":{"key":"description","value":"update"}}'

4. 删除数据

   + 删除一条

     curl -X DELETE localhost:7460/operation/delete -d '{"key":"name","value":"op"}'









