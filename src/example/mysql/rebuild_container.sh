echo ">>> build exec file"
docker run --rm -v /root/own/src:/go/src -v /root/own/bin:/go/bin golang:1.11.5-alpine3.9 sh -c "export GOPATH=/go && cd /go/src/example/mysql && go build -o mysql_basic"
echo ">>> go build completed"
docker rm -f mysql_basic
docker build -t mysql_basic /root/own/src/example/mysql
echo ">>> docker build completed"
docker run -it -d \
  --net host \
  --name mysql_basic \
  -e MYSQL_ROOT_PASSWORD=123456 \
  mysql_basic
