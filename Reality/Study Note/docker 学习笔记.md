### docker 学习笔记

1. 修改docker镜像的tag

   `docker tag [imageID] REPOSITORY:TAG`

   例：`docker tag 96c120fadccf portal:latest`

2. image 导出为 tar 文件

   `docker save [image:tag] > [image.tar]`

   `docker save [image:tag] -o [image.tar]`

3. 导入image.tar

   `docker load < [image.tar]`

   `docker load -i [image.tar]`

4. 通过dockerfile制作image

   `docker build -t elasticsearch docker/elasticsearch`

   >  `-t` ，添加image标签
   >
   >  最后 `docker/elasticsearch` ，指定 dockerfile 的文件位置
   >
   >  “ . ” 表示当前路径

5. 制作 docker image 之后，通过 docker run 命令运行容器

6. 删除无用docker镜像

   `docker rmi $(docker images | grep none | awk '{print $3}')`

   + $(xxxx)：将xxxx作为参数，传入docker rmi
   + awk '{print $3}'：输出第三列

7. docker中 -i 参数和 -t 参数

   -t，表示分配一个伪终端
   
   -i，表示可以交互，可以从键盘输入（stdin），并且获得反馈
   
   
   
   docker exec container sh
   
   无返回
   
   docker exec -t container sh
   
   有一个可交互的终端，但是输入无响应
   
   docker exec -i container sh
   
   没有终端，但是输入有响应，没有格式
   
   docker exec -it container sh
   
   与正常的终端交互无差别
   
   
   
   [参考链接](https://blog.csdn.net/claram/article/details/104228727)
   
   
   
8. 获取某行某列

   ![image-20200120113157165](C:\Users\Jarvis.LAPTOP-HV4II8QE\AppData\Roaming\Typora\typora-user-images\image-20200120113157165.png)

   
