pkill xingtu

sleep 1s

go build -o /data/balfour/zhiyuxingtu/build/xingtu /data/balfour/zhiyuxingtu/main.go

sleep 1s

nohup ./build/xingtu > ./logs/run.log 2>&1 &

tail -f ./logs/run.log
