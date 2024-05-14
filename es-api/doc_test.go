package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var (
	EsClient *elasticsearch.Client
	once     sync.Once
)

// GetredisClient 返回 Redis 客户端实例
func GetEsClient() *elasticsearch.Client {
	once.Do(func() {
		var err error
		EsClient, err = elasticsearch.NewClient(
			elasticsearch.Config{
				Addresses: []string{"http://175.178.2.100:9200"},
				Username:  "",
				Password:  "",
			},
		)
		if err != nil {
			fmt.Println("err", err.Error())
		}
	})
	return EsClient
}

func TestDocIndex(t *testing.T) {
	es := GetEsClient()

	// for i := 0; i < 20; i++ {
	// 	doc := map[string]interface{}{
	// 		"id":            i,
	// 		"supplier_name": "supplier_name",
	// 		"type_id":       i,
	// 		"type_name":     "type_name",
	// 		"category_id":   i,
	// 		"category_name": "category_name",
	// 		"brand_name":    "dio",
	// 		"goods_price":   i,
	// 	}
	// 	jsonDoc, err := json.Marshal(doc)
	// 	if err != nil {
	// 		fmt.Printf("Error encoding document: %s \n", err.Error())
	// 	}
	// 	req := esapi.IndexRequest{
	// 		Index:      "supore_mall",
	// 		DocumentID: strconv.Itoa(i),
	// 		Body:       strings.NewReader(string(jsonDoc)),
	// 		Refresh:    "true",
	// 	}
	// 	// 发送文档请求
	// 	res, err := req.Do(context.Background(), es)
	// 	if err != nil {
	// 		log.Fatalf("Error indexing document: %s", err)
	// 	}
	// 	fmt.Printf("response: %s \n", res.String())
	// 	defer res.Body.Close()
	// }

	// 01. 创建文档
	// doc := map[string]interface{}{
	// 	"title":   "Golang Elasticsearch Example",
	// 	"content": "This is a simple example of adding a document to Elasticsearch using Golang.",
	// 	"price":   12.00,
	// }
	//将文档结构编码为JSON
	// jsonDoc, err := json.Marshal(doc)
	// if err != nil {
	// 	fmt.Printf("Error encoding document: %s \n", err.Error())
	// }
	// // 构建文档请求
	// req := esapi.IndexRequest{
	// 	Index:      "supore_mall",
	// 	DocumentID: "1", // 可选，如果未提供，Elasticsearch将为文档分配一个唯一的ID
	// 	Body:       strings.NewReader(string(jsonDoc)),
	// 	Refresh:    "true",
	// }
	// // 发送文档请求
	// res, err := req.Do(context.Background(), es)
	// if err != nil {
	// 	log.Fatalf("Error indexing document: %s", err)
	// }
	// defer res.Body.Close()

	// // 检查响应状态码
	// if res.IsError() {
	// 	log.Fatalf("Error response received: %s", res.Status())
	// }
	// log.Println("Document indexed successfully")

	// 02. 删除文档数据
	// req := esapi.DeleteRequest{
	// 	Index:      "supore_mall",
	// 	DocumentID: "1",
	// }

	// // 发送删除文档请求
	// res, err := req.Do(context.Background(), es)
	// if err != nil {
	// 	log.Fatalf("Error deleting document: %s", err)
	// }
	// defer res.Body.Close()

	// 03 查询文档数据
	//构建查询DSL
	// query := `
	// {
	// 	"query": {
	// 		"term": {
	// 			"goods_price": 23
	// 		}
	// 	}
	// }` // 查询价格等于23

	// query1 := `
	// {
	// 	"query": {
	// 		"range": {
	// 			"goods_price": {
	// 				"gt": 14
	// 			}
	// 		}
	// 	}
	// }` // 查询价大于14
	// }`
	// 构建查询DSL
	query := `
		{
			"query": {
				"match_all": {}
			},
			"from": "4",
			"size": "10"
		}`

	// 创建搜索请求
	req := esapi.SearchRequest{
		Index: []string{"supore_mall"}, // 索引名称
		Body:  bytes.NewReader([]byte(query)),
	}

	// 执行搜索请求
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error executing search: %s", err)
	}
	defer res.Body.Close()

	// 检查响应状态
	if res.IsError() {
		log.Fatalf("Error response received: %s", res.Status())
	}

	// 读取响应体
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s", err)
	}

	// 解析搜索结果
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// for k, v := range result {
	// 	fmt.Printf("%s -> %+v\n", k, v)
	// }

	// fmt.Printf("%s -> %+v\n", v)
	// 获取hits字段
	hits, found := result["hits"].(map[string]interface{})
	if !found {
		log.Fatalf("Hits not found in response")
	}

	// 获取hits中的hits数组
	hitsArray, found := hits["hits"].([]interface{})
	if !found {
		log.Fatalf("Hits array not found in response")
	}

	// 遍历hits数组，提取每个文档的信息
	for _, hit := range hitsArray {
		hitMap, found := hit.(map[string]interface{})
		if !found {
			log.Fatalf("Hit is not a map")
		}

		// 获取文档信息
		source, found := hitMap["_source"].(map[string]interface{})
		if !found {
			log.Fatalf("Source not found in hit")
		}
		fmt.Println("Document Info:", source)
		// 输出文档信息
		fmt.Printf("Document ID: %v\n", source["id"])
		fmt.Printf("Document Index: %v\n", hitMap["_index"])
		fmt.Printf("Document Score: %v\n", hitMap["_score"])
	}

}
