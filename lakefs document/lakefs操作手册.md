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

   ` docker run -d -p 27017:27017 -v /root/mnt/:/data/db --name mongo mongo`

5. 执行lakefs，测试s3

   `./lakefs --tier-type=das3 --mount-point /root/mnt/ --fast-root-path=/root/fast/ --cleanup-interval=10 --transit-interval=10 --admin-listen-addr=http://10.2.174.54:5000 --s3-endpoint http://123.177.21.80:8004 --s3-ak system --s3-sk 123456 --s3-bucket a12345 --debug 1`
   
6. 指定文件，以及行为和时间

   `./set.sh /root/mnt/1.txt migrate 10`

7. 查看执行lakefs后输出的log

   

8. 进入mongo查看inode信息





#### Q & A

1. mongo中md5信息没有写入，state为0

   ---

2. 通过makefile，执行make lakefs，报错无法下载golang，终止退出

   ---

3. ls 显示 `ls: cannot access mnt: Transport endpoint is not connected`；mnt

   fusermount -uz mnt，强制unmount mnt目录

   ---

4. mongo容器启动失败，总处于created状态

   解决 **3** 之后，删除容器再重启，不再出现该问题

   ---

5. 执行lakefs后，显示日志

   > 2019/04/24 07:07:03  000e transition.go:97 ▶ INFO 064 start cleaning up backed files
   > 2019/04/24 07:07:03  000e transition.go:291 ▶ INFO 065 finished cleanup 0 files
   > 2019/04/24 07:07:03  000d transition.go:68 ▶ INFO 066 start transiting expired files
   > 2019/04/24 07:07:03  0075 transition.go:268 ▶ INFO 067 finished transit 1 files
   > 2019/04/24 07:07:03  0044 transition.go:113 ▶ DEBU 068 Locked inode 19126341 with lockID 83417a2c-a731-458a-8efa-3e1bb5fa2539
   > 2019/04/24 07:07:03  0044 transition.go:130 ▶ DEBU 069 migrateFile inode 0xc0001d4ab0 &{Inode:19126341 CreateDate:2019-04-23 05:31:27.162 -0400 EDT ExpireDate:2019-04-24 07:06:35.154 -0400 EDT ExpireAction:migrate State:0 Size:0 Md5Sum: RestoreDate:<nil> BackupDate:<nil> updateDate:{wall:0 ext:0 loc:<nil>} fs:0xc000121ce0}
   > 2019/04/24 07:07:03  0044 das3.go:39 ▶ DEBU 06a Put object a123456/123d845.20190424.070703.862
   > 2019/04/24 07:07:03  005d fs.go:792 ▶ DEBU 06b read 4 bytes from /root/fast/1.txt @ op.Offset 0
   > 2019/04/24 07:07:03  005d fs.go:792 ▶ DEBU 06c read 0 bytes from /root/fast/1.txt @ op.Offset 4
   > 2019/04/24 07:07:03  0044 das3.go:43 ▶ DEBU 06d Put a123456/123d845.20190424.070703.862 got **The specified bucket does not exist.**
> 2019/04/24 07:07:03  0044 inode.go:190 ▶ DEBU 06e backed up inode 19126341, 0 bytes md5sum ba1f2511fc30423bdbb183fe33f3dd0f
   > 2019/04/24 07:07:03  0044 transition.go:119 ▶ DEBU 06f Unlocked lockID 83417a2c-a731-458a-8efa-3e1bb5fa2539

   日志中显示

   Put a123456/123d845.20190424.070703.862 got **The specified bucket does not exist；指定桶不存在

   finished transit 1 files；transit完成

   backed up inode 19126341, 0 bytes md5sum d41d8cd98f00b204e9800998ecf8427e；md5已经计算完毕

   进入mongo，查询只能发现inode，没有md5，state，restoredate，backupdate

   

   **尝试解决：**

   ​	明天（4.25）

   1. 确认参数全部都有，并且无误

   2. 创建新的bucket，用新的试试

   3. 查看log对应位置的源码

      4.27
   
   4. 问超哥

      **更换超哥的minio节点，上传成功，初步预计是大连的问题**

      **全部问题已经解决，和大连对接s3接口成功**

      