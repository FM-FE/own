# Mongo加密部署和使用教程

## 部署

![image-20200303150349972](C:\Users\Friday\AppData\Roaming\Typora\typora-user-images\image-20200303150349972.png)

### 准备证书

1. openssl genrsa -aes256 -passout pass:123456 -out server.key 1024

2. openssl req -new -x509 -days 3650 -passin pass:123456 -key server.key -out server.crt \
   -subj "/C=CN/ST=Beijing/L=Beijing/O=hello/OU=dev/CN=hello.com/emailAddress=self@hello.com"

3. cat server.crt server.key > server.pem

### 下载启动mongo

1. docker pull mongo:4.2

2. docker run --name mongo -p 27017:27017 -d --restart=always --log-opt max-size=50m \
   -v **~/mongo/db:/data/db** -v /etc/localtime:/etc/localtime:ro -v ~**/mongo/ssl/**:/etc/tls/ \
   mongo:4.2 --tlsMode requireTLS --tlsDisabledProtocols TLS1_0,TLS1_1 \
   --tlsCertificateKeyFile /etc/tls/server.pem --tlsCertificateKeyFilePassword 123456 --auth
   
   > 注意，上图加粗部分需要视情况替换
   >
   > 下面的参数不需要替换
   >
   > --tlsCertificateKeyFile /etc/tls/server.pem

### 添加nasuser用户

1. docker exec mongo bin/bash -c "echo 'admin = db.getSiblingDB(\"admin\"); admin.createUser({user: \"nasuser\" , pwd: \"0mkjwurNA#Lat7V4\", roles: [{ role:\"root\", db: \"admin\"},\"readWriteAnyDatabase\"] })' | mongo --tls --tlsCertificateKeyFile "/etc/tls/server.pem" --tlsAllowInvalidCertificates --tlsCertificateKeyFilePassword 123456 -u nasuser -p 0mkjwurNA#Lat7V4

