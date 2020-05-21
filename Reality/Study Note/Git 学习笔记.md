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
   
   
   
4. git config --global core.autocrlf ture

   `git add .` 遇到报错信息：

   `warning: LF will be replaced by CRLF in .gitignore.
   The file will have its original line endings in your working directory.`

   Windows系统下，换行符为CRLF

   Linux或Mac系统下，换行符为LF

   在多人协作情况下，Windows用户提交的文件，由于换行符不同，会被Linux/Mac用户视为修改了整个文件，没有办法使用 git diff 。

   所以需要设置 core.autocrlf 

   ```
   # 提交时转换为LF，检出时转换为CRLF
   git config --global core.autocrlf true
   
   # 提交时转换为LF，检出时不转换
   git config --global core.autocrlf input
   
   # 提交检出均不转换
   git config --global core.autocrlf false
   ```

   所以，在Windows系统下，应该设置 core.autocrlf 为 true ，在提交时统一换行符为 LF ，在检出时，调整为Windows的换行符 CRLF

   

5. Git 上传文件的大小限制为 100M

   如果需要上传超过 100M 的文件，请参考链接

   [个人博客 Harrlet Land][https://harttle.land/2016/03/22/purge-large-files-in-gitrepo.html]

   [简书](<https://www.jianshu.com/p/f4f34c67707a>)

   

6. gitignore

   Windows系统，在 git bash中，touch .gitignore

   ~~不能新建文本文件，再重命名文件；Windows默认gitignore为文件后缀，要求必须键入文件名~~

   用 notepad++ 编辑 .gitignore ，在其中添加将要忽略的文件名。

   **注意：要在提交之前，做好 gitignore；如果已经提交了要忽略的文件，备份忽略文件后，`git reset --hard <log ID>`，回到提交之前**

   写 gitignore 的文件规则

   "/" 开头表示，仅匹配当前目录下的文件

   空白开头表示，匹配当前目录以及子目录下的所有文件

   如，`deploy/` 可以匹配当前目录下的 `deploy/`，也可以匹配到 `src/deploy/`

   而，`/deploy/` 仅可以匹配到当前目录下的 `deploy/`

   

   

7. git branch

   显示本地分支

   

8. git checkout [-b] 本地分支名

   切换到本地分支

   加 -b 参数，创建分支并切换







### 远程

1. git push origin master

   将本地仓库提交到远程仓库的master分支
   
   
   
2. git push -f origin master

   当本地分支和远程分支不匹配时，强制推送，更新远程仓库
   
   
   
3. git remote add origin [git@IP ADDRESS:REPO.git]

   添加远程仓库

   如果希望使用 git clone，需要把本地rsa公钥复制到Git

   [添加远程仓库参考链接](<https://www.liaoxuefeng.com/wiki/896043488029600/898732864121440>)

   [复制公钥参考链接](<https://www.liaoxuefeng.com/wiki/896043488029600/896954117292416>)

   

4. git pull origin master --allow-unrelated-histories

   合并两个不相干的库，本地库和远程库

   

5. git checkout -b 本地分支名 origin/远程分支名

   切换到远程分支，并且在本地同步创建分支

   

6. [将本地代码上传到远程的几种情景](https://blog.csdn.net/programmer_at/article/details/78011705)

   

7. `git push origin master(本地分支):v2.0(远程分支)`

   将本地的master分支推送到远程的v2.0分支

   

8. `git pull origin v2.0(远程分支):master(本地分支)`

   将远程的v2.0分支下载 ( fetch ) 到本地master分支 ( merge )  









### 处理冲突

1. stash

   git statsh 

   git pull

   git stash pop

   

2. merge

   git pull

   决定保留和更改

   重新commit，push





### 强制回退到本地或远程的某个commit

1. `git log` 或 `git log origin/master` 查看commit id
2. `git reset --hard “commit id”`



### git误删恢复

1. git add之后，还未commit，就使用reset --hard命令，回退到add之前，使得暂存区修改失效

   [参考链接](https://www.cnblogs.com/hope-markup/p/6683522.html)

   1. 根目录下，输入找回命令：`git fsck --lost-found`

      ![image-20200521153742415](C:\Users\Jarvis.LAPTOP-HV4II8QE\AppData\Roaming\Typora\typora-user-images\image-20200521153742415.png)

   2. 等待命令执行结束，用ide打开 .git/lost-found/other 目录，找到被误删的文件

      

      

       

      

      

   











