package main

import (
	"ArSearch/pkg/service"
	"ArSearch/pkg/service/service_schema"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

func main() {
	kafkaURL := "localhost:9092"
	//kafkaURL := "45.76.151.181:9092"
	topic := "quickstart-events"
	groupID := "groupId"
	reader := service.GetKafkaReader(kafkaURL, topic, groupID)
	wg := sync.WaitGroup{}
	for {
		wg.Add(1)

		go func() {

			m, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Fatalln(err)
			}

			v := service_schema.ArData{}
			json.Unmarshal(m.Value, &v)

			link1 := fmt.Sprintf("htts:mirror.xyz/%s/%s", v.Digest, v.OriginalDigest)
			mirroData := service_schema.MirrorData{
				Id:             m.Offset,
				Title:          v.Content.Title,
				Content:        v.Content.Body,
				Digest:         v.Digest,
				Link:           link1,
				OriginalDigest: v.OriginalDigest,
				ArweaveTx:      v.Digest,
			}
			data, err1 := service.SaveMirrorData(&mirroData)
			if err1 != nil {
				fmt.Println("err===>", err1)
			}
			fmt.Println("data=====>", data)

			wg.Done()
		}()

		//fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		//time.Sleep(time.Second * 50)
	}
	wg.Wait()
}
