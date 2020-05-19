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

   

4. 完成NAS V2.0的修改

   **done**

---

## 20200507

1. 节点列表过滤node测试

   **done**

2. 执行deploy之后，重启lakefs的docker

   + golang版本

     init强于main函数的原因

     

     1. 判断是否有升级标志文件（/tank/web/updated）
     2. 如果存在updated文件，表示已经升级，不用做任何操作
     3. 如果不存在updated文件，表示还未升级
        1. curl命令请求，获得需要操作的容器
        2. 停止正在运行的容器
        3. 删除已停止的容器
        4. 删除旧镜像
        5. 载入新镜像
        6. 运行新容器

   + python版本

---

## 20200508

1. 测试init是否成功
   + 尝试python版本
2. 在删除挂载时候，判断对应卷的lakefs-fuse容器是否存在

---

## 20200509

1. 简化dockerfile和启动脚本

---

## 20200514

在下午三点之前解决除了，数据库相关的其他事情

有事就要尽快做

1. v2.0版本将debug级别的日志调整为warnning级别

   **done**

2. v3.0版本，挂载卷，如果已有容器，报错信息调整

   **done**

3. v3.0版本，添加disk新接口转发

   **done**

4. v3.0版本，删除卷时候，删除数据库

---

## 20200518

1. lakefs-transition新接口

   

2. 删除卷时，删除对应数据库

   

3. 修改nas/list/all的接口

   **done**

---

## 20200519

活都好做，什么时候做完呢

1. lakefs-transition新接口

   

2. 删除卷时，删除对应数据库

   

3. 挂载外部nas迁移策略没有选项

   直接挂载似乎不可选择

   后续关联光盘组可以选择

   

4. 新建盘匣组 >> 光盘库重名

   查看gdas文档，联系王钊解决问题













