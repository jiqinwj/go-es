# go-es
go es   操作
## 启动kb 服务
```
docker run --init -d --name kb \
     -e elasticsearch.hosts="http://127.0.0.1:9200"   \
     -p 5601:5601  \
     kb:74
```
