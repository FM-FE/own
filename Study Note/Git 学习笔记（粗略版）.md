### Git 学习笔记

1. git是版本控制工具
2. git的使用分远程和本地

### 本地

1. git init

   初始化当前目录为 git 仓库

2. git add [ filename ] or **.**

   添加文件到stage，如果写 **.** 则表示添加当前所有文件到stage

3. git commit -m " note "

   提交修改后的文件

### 远程

1. git push origin master

   将本地仓库提交到远程仓库的master分支
   
2. git push -f origin master

   当本地分支和远程分支不匹配时，强制推送，更新远程仓库