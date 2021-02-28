# redis

## 遇到的问题

1. 程序无法编译，由于缺少包

   查GitHub官方资料，提示需要

   ```
   go mod init github.com/my/repo
   go get github.com/go-redis/redis/v8
   ```

   但是，本地并不能go get，好像是因为墙的原因

   go mod是什么，怎么用也不清楚

