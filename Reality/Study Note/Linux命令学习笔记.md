## Linux命令学习笔记

1. 压缩解压命令(tar)

   + 压缩命令 
   
     `tar cvzf xxx.tar.gz /root/dir`
   
   + 解压命令
   
     `tar xvzf xxx.tar.gz -C /root/dir`
   
   c： 创建归档
   
   x ：取消归档
   
   z：压缩或解压 
   
   f：指定要操作的文件名
   
   > tar文件只归档，不压缩
   >
   > tar.gz文件，以gzip形式压缩，需要在命令中添加 “z” 参数 