# go-es
go es   操作
## 启动kb 服务
```
docker run --init -d --name kb \
     -e elasticsearch.hosts="http://192.168.1.20:9200"   \
     -p 5601:5601  \
     kb:74
```
