## dam创建S3存储服务

**方一**

1. 创建命令

   `curl -X PUT --user admin:test1234 localhost:7090/users -d '{"uid":"fangyi","pass":"123123","vaulttag":"fangyi"}'`

2. 查询ak，sk

    `curl -X GET --user admin:test1234 localhost:7090/users`

   返回ak，sk

   `[{"uid":"fangyi","pass":"123123","ak":"mhze4v8mW0agzjqI","sk":"cJIiV8Kh1vn1uRxdMRLHnnvsy3SDzj4b","vaulttag":"fangyi"}]`

3. 创建存储桶

   拷贝aws-cli的安装文件到服务器，用aws命令创建存储桶

   1. 设置ak，sk

      进入 /root/.local/lib/aws/bin

      执行 `./aws configure`

      输入ak，sk

   2. 创建桶

      `aws --endpoint-url=http://IP:9000 s3api create-bucket --bucket BucketName`

      











