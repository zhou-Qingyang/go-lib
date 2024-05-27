package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
)

//Writer中的Topic和Message中的Topic是互斥的，同一时刻有且只能设置一处。

func TestXxx(t *testing.T) {
	// writeByConn 基于Conn发送消息
	topic := "my-topic"
	partition := 0

	// 连接至Kafka集群的Leader节点
	conn, err := kafka.DialLeader(context.Background(), "tcp", "175.178.2.100:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	// 设置发送消息的超时时间
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	// 发送消息
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	// 关闭连接
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func TestCreateTopic(t *testing.T) {
	// 指定要创建的topic名称
	topic := "my-topic2"
	// 连接至任意kafka节点
	conn, err := kafka.Dial("tcp", "175.178.2.100:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// 获取当前控制节点信息
	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	// 连接至leader节点
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	// 创建topic
	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}

func TestGetTopicList(t *testing.T) {
	conn, err := kafka.Dial("tcp", "175.178.2.100:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}
	// 遍历所有分区取topic
	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}

func TestReader(t *testing.T) {
	// 创建Reader
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"223.82.117.13:9092"},
		Topic:          "VisionQrCodeRecord12",
		Partition:      0,
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
	})
	r.SetOffset(0) // 设置Offset
	// 接收消息
	defer func() {
		// 程序退出前关闭Reader
		if err := r.Close(); err != nil {
			log.Fatal("failed to close reader:", err)
		}
	}()
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout reached. Exiting program.")
			return
		default:
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				break
			}
			fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
		}
	}
}
func TestReader2(t *testing.T) {
	// 创建Reader
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"223.82.117.13:9092"},
		Topic:     "VisionQrCodeRecord12",
		Partition: 0,
		GroupID:   "your-consumer-group-id1", // 设置消费者组ID
		MaxBytes:  10e6,                      // 10MB
	})
	r.SetOffset(0) // 设置Offset
	// 接收消息
	defer func() {
		// 程序退出前关闭Reader
		if err := r.Close(); err != nil {
			log.Fatal("failed to close reader:", err)
		}
	}()
	ctx := context.Background()
	for {
		// 获取消息
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		// 处理消息
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		// 显式提交
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}

func TestWriteToOneTopic(t *testing.T) {
	// 创建一个writer 向topic-A发送消息
	// 必须先存在一个topic 要么就带上 AllowAutoTopicCreation这个 字段
	w := &kafka.Writer{
		Addr:         kafka.TCP("175.178.2.100:9092"),
		Topic:        "my-topic",
		Balancer:     &kafka.LeastBytes{}, // 指定分区的balancer模式为最小字节分布
		RequiredAcks: kafka.RequireAll,    // ack模式
		Async:        true,                // 异步
		// AllowAutoTopicCreation: true,  // 自动创建topic
	}
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func TestWriteToManyTopic(t *testing.T) {
	w := &kafka.Writer{
		Addr: kafka.TCP("175.178.2.100:9092"),
		// 注意: 当此处不设置Topic时,后续的每条消息都需要指定Topic
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	err := w.WriteMessages(context.Background(),
		// 注意: 每条消息都需要指定一个 Topic, 否则就会报错
		kafka.Message{
			Topic: "topic-A",
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Topic: "topic-B",
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Topic: "topic-C",
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
