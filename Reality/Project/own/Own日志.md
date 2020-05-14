# Own实现日志

## 20200428

1. 厘清后端究竟能做到哪一步，究竟能不能做到可用

## 20200429

1. 完成mongo的按条件增删改查的restful操作

> 关于own还能做什么
>
> 比起实现具体的功能更重要的是用上各种技术
>
> 是能熟练的用上各种技术
>
> 能熟练的用上各种技术的基本操作
>
> 目前，已经在用的技术有
>
> 1. docker
>
>    + swarm
>    + service
>
> 2. go
>
>    + 存储树形结构
>
> 3. mongo
>
>    + 增删改查里，只实现了增
>
>      还有删改查，怎么结合条件删改查
>
>      像sql一样
>
>    + 给数据加索引
>
>    + 加密mongo
>
>    + mongo用户管理
>
>    + mongo集群
>
> 还需要的准备的技术
>
> 1. redis
> 2. kafka
>
> 
>
> mongo的增删改查熟练的实现了，就算是初步准备好了
>
> 也就是说，5月2号之前
>
> 最起码，要实现熟练的mongo操作

## 20200504

1. 完成mongo的基本操作
   + **mongo包的例子** 当中有关于查询条件的具体例子

---

## 20200505

1. 完成mongo基本操作
   + mongo的删除和修改操作

---

## 20200507

> Own的作用就是铺路
>
> 给工作中的新技术铺路
>
> 给跳槽路上需要的技术铺路

---

## 20200508

找到了[beego教程](https://github.com/astaxie/build-web-application-with-golang)，里面提到了[如何通过golang使用MySQL](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.2.md)

还找到了[gin教程]()，项目中用到了redis

1. MySQL

   1. 通过docker compose部署MySQL

      [部署文档](O:\own\doc\deploy\mysql\MySQL部署教程(docker compose).md)

   2. 通过[go-sql-driver](https://github.com/go-sql-driver/mysql )操作MySQL

   3. 回顾大学学习的sql语句，完成基本操作


---

## 20200514

1. docker compose学习，留下部署文档

2. mysql

   docker compose部署mysql，并留下学习文档

   golang对mysql的使用，留下学习文档  [golang使用mysql](直接连接到代码)

   **在部署文档后，直接附上代码链接**

3. mongo

   docker compose部署加密mongo，留下文档

   golang对加密mongo的使用 [golang使用加密mongo](代码链接)

   **在部署文档后，直接附上代码链接**

   





