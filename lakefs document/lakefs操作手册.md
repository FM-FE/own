### lakefs操作手册

可能需要连接外网

1. 安装docker

   在线安装或离线安装

   [离线安装参考](https://blog.csdn.net/corbin_zhang/article/details/81325114)

2. 建立dambuild

   复制newdam代码到本地

   执行make dambuild

3. lakefs目录下，找到Makefile，执行make lakefs

   `docker build -t lakefs src/ehualu.com/lakefs`

   生成docker镜像

   `go build -o /go/src/ehualu.com/lakefs/lakefs ehualu.com/lakefs`

   在src/ehualu.com/lakefs目录下，生成lakefs可执行文件

4. 准备mongo镜像，运行mongo容器

   ` docker run -d -p 27017:27017 -v /root/mongo/:/data/db --name mongo mongo`

5. 执行lakefs，测试s3

   `./lakefs --tier-type=das3 --mount-point /root/mnt/ --fast-root-path=/root/fast/ --cleanup-interval=10 --transit-interval=10 --admin-listen-addr=http://10.2.174.121:5000 --s3-endpoint http://123.177.21.80:8004 --s3-ak system --s3-sk 123456 --s3-bucket das3test --debug 1`
   
6. 指定文件，以及行为和时间

   `./set.sh /root/mnt/1.txt migrate 10`

7. 查看执行lakefs后输出的log

   

8. 进入mongo查看inode信息





#### Q & A

1. mongo中md5信息没有写入，state为0

   ---

3. ls 显示 `ls: cannot access mnt: Transport endpoint is not connected`；mnt

   > fusermount -uz mnt，强制unmount mnt目录

   ---

3. mongo容器启动失败，总处于created状态

   解决 **2** 之后，删除容器再重启，不再出现该问题

   ---

4. set 历史同名文件报错

   `./set.sh /root/mnt/1.txt migtate 10`

   报错

   `setfattr: /root/mnt/1.txt: No such file or directory`
   `setfattr: /root/mnt/1.txt: No such file or directory`

   > 之前 fast 路径下，有过1.txt
   >
   > 也对1.txt执行过set脚本
   >
   > 删除1.txt后，新建同名文件
   >
   > 再执行set脚本，报错如上
   >
   > 怀疑不能创建历史同名文件