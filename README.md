ArSearch

```
kafka 下载
https://www.apache.org/dyn/closer.cgi?path=/kafka/3.2.0/kafka_2.13-3.2.0.tgz

数据sync 

1.确保kafka 已启动
nohup bin/zookeeper-server-start.sh config/zookeeper.properties >> zk.log &
nohup bin/kafka-server-start.sh config/server.properties >> kafka-server.log & 

#仅第一次无topic 时需指定topic 
bin/kafka-topics.sh --create --topic arsearch-topic --bootstrap-server localhost:9092

#查询线上topic 内容(可以在写入kafka 时读一下看看非必须)
bin/kafka-console-consumer.sh --topic arsearch-topic --from-beginning --bootstrap-server localhost:9092

2.sync 数据至线上es 中

#默认写入线上es http://45.76.151.181:9200
#仅需测试 可以代码调整成localhost
-nohup go run sync/sync2.go >> sync2kafka.log &
-nohup go run colly/main4.go >> sync2es.log & 


3.开启线上服务
nohup go run main.go >> search.log & 

测试是否可用
curl http://45.76.151.181:9999/query_mirror?q=[search keyword]

```