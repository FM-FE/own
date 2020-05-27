# docker部署MySQL教程

## 简介

MySQL是一种关系型数据库

## 下载

`docker pull mysql[:tag]`

用tag指定版本，默认为latest

[docker hub官方镜像参考](https://hub.docker.com/_/mysql)

## 运行容器

`docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql`









## 问题及解决

1. docker pull命令失败

   可能是因为国内防火墙的原因

   方法一：采用国内镜像

   [国内镜像站](https://hub.daocloud.io/repos)

   `docker pull daocloud.io/library/mysql`

   [参考链接](https://blog.csdn.net/w_bu_neng_ku/article/details/78765251)

   或者

   修改`/etc/docker/daemon.json`文件（如果没有就直接创建）

   ```json
   {
    "registry-mirrors": ["https://docker.mirrors.ustc.edu.cn/"]
   }
   ```

   `systemctl daemon-reload`

   `systemctl restart docker`

   `docker pull mysql`

   [参考链接](https://blog.csdn.net/qq_39329616/article/details/89640731)

   还有[其他镜像站](https://blog.csdn.net/qq_39329616/article/details/89640731)可选择，这种方法不能保证获取的镜像是最新的

   

   方法二：用一台可以科学上网的电脑代理docker的http/https请求

   [官方链接](https://docs.docker.com/config/daemon/systemd/)

   [参考链接](https://blog.csdn.net/qq_42684642/article/details/85302222)

   

   













