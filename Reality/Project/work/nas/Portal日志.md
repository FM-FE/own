# Portal实现日志

## 20200428

1. 获取所有节点的可用bricks

   + go routine和channel
   + 

   ---

2. 挂载全部节点的gluster卷

   + go routine和channel
   + 

---

## 20200430

组合挂载情况和transition

1. 在挂载之前先执行启动卷操作，不做返回判断

   **DONE**

2. 节点列表显示时，过滤node的问题

   目前10.2.174.114也在集群中，但是portal代理时按照node过滤了

   

3. 删除挂载

   1. 删除lakefs
   2. 删除gluster挂载

   **DONE**

---

## 20200506

1. 节点列表显示时，过滤node的问题

   目前10.2.174.114也在集群中，但是portal代理时按照node过滤了

   关键的问题是在解析hosts文件时，不能按照node解析

   

2. 前端批量更改策略报错

   **done**

3. 执行deploy之后，重启lakefs的docker

   











