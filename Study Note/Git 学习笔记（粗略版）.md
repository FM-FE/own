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

