ArSearch

```
数据sync 

1.确保kafka 已启动
nohup bin/zookeeper-server-start.sh config/zookeeper.properties >> zk.log &
nohup bin/kafka-server-start.sh config/server.properties >> kafka-server.log & 

#仅第一次无topic 时需指定topic 
bin/kafka-topics.sh --create --topic arsearch-topic --bootstrap-server localhost:9092

#查询线上topic 内容
bin/kafka-console-consumer.sh --topic arsearch-topic --from-beginning --bootstrap-server localhost:9092

```