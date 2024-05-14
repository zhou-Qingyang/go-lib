package main

import (
	"context"
	"log"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func TestMapping(t *testing.T) {
	es := GetEsClient()
	// 01 删除mapping
	// req := esapi.IndicesDeleteRequest{
	// 	Index: []string{"supore_mall"},
	// }
	// // 发送创建索引请求
	// res, err := req.Do(context.Background(), es)
	// if err != nil {
	// 	log.Fatalf("Error creating index: %s", err)
	// }
	// defer res.Body.Close()

	// 02 添加mapping映射
	mapping := `{
		"mappings": {
			"properties": {
				"supplier_name": {
					"type": "text"
				},
				"brand_name": {
					"type": "text"
				},
				"category_id": {
					"type": "long"
				},
				"category_name": {
					"type": "text"
				},
				"type_name": {
					"type": "text"
				},
				"type_id": {
					"type": "long"
				},
				"goods_price": {
					"type": "long"
				},
				"goods_img": {
					"type": "text"
				},
				"id": {
					"type": "long"
				}
			}
		}
	}`

	// 创建索引请求
	req := esapi.IndicesCreateRequest{
		Index: "supore_mall",
		Body:  strings.NewReader(mapping),
	}

	// 发送创建索引请求
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error creating index: %s", err)
	}
	defer res.Body.Close()
	// 检查响应状态
	if res.IsError() {
		log.Fatalf("Error response received: %s", res.Status())
	}

	log.Println("Index created successfully with mapping.")

}
