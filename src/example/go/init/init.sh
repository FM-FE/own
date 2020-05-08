echo ">>> build exec file"
docker run --rm -v /root/own/src:/go/src -v /root/own/bin:/go/bin golang:1.11.5-alpine3.9 sh -c "export GOPATH=/go && cd /go/src/example/go/init && go build -o alpine_init"
echo ">>> go build completed"
docker rm -f init
docker build -t init /root/own/src/example/go/init
echo ">>> docker build completed"
docker run -it -d \
--net host \
--name init \
init