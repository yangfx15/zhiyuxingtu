pkill xingtu

sleep 1s

go build -o /data/balfour/zhiyuxingtu/xingtu /data/balfour/zhiyuxingtu/main.go

sleep 1s

nohup ./xingtu > run.log 2>&1 &

tail -f run.log
