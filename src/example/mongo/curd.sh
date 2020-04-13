docker run --rm -v /root/own/src:/go/src -v /root/own/bin:/go/bin golang:1.11.5-alpine3.9 sh -c "export GOPATH=/go && cd /go/src/example/mongo && go build -o curd"
echo ">>> go build completed"
docker rm -f curd
docker build -t curd /root/own/src/example/mongo
echo ">>> docker build completed"
docker run -it -d \
--net host \
--name curd \
curd