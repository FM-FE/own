# Question List

1. 如何让github通过ssh协议下载

   `git clone git@github.com:mongodb/mongo-go-driver.git`

   命令报错

   ```
   $ git clone git@github.com:mongodb/mongo-go-driver.git
   Cloning into 'mongo-go-driver'...
   Connection reset by 13.250.177.223 port 22
   fatal: Could not read from remote repository.
   
   Please make sure you have the correct access rights
   and the repository exists.
   ```

   

2. git clone https 下载失败

   ```
   $ git clone https://github.com/mongodb/mongo-go-driver.git
   Cloning into 'mongo-go-driver'...
   remote: Enumerating objects: 30, done.
   remote: Counting objects: 100% (30/30), done.
   remote: Compressing objects: 100% (27/27), done.
   error: RPC failed; curl 18 transfer closed with outstanding read data remaining
   fatal: The remote end hung up unexpectedly
   fatal: early EOF
   fatal: index-pack failed
   ```

   

3. context.WithTimeout((context.Background(), 10*time.Second)) 是否表示10秒后，连接过期

   https://juejin.im/post/5d4c3e64e51d4561d044cc8a







