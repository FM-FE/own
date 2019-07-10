### docker 学习笔记

1. image 导出为 tar 文件

   `docker save [image:tag] > [image.tar]`

   `docker save [image:tag] -o [image.tar]`

2. 导入image.tar

   `docker load < [image.tar]`

   `docker load -i [image.tar]`
   
3. 通过dockerfile制作image

   `docker build -t elasticsearch docker/elasticsearch`

   >  `-t` ，添加image标签
   >
   > 最后 `docker/elasticsearch` ，指定 dockerfile 的文件位置

4. 制作 docker image 之后，通过 docker run 命令运行容器