package main

import (
	"encoding/json"
	"fmt"
)

// 定义 JSON 对应的结构体
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

func main() {
	jsonStr := `{"title": "Golang入门", "author": "张三", "price": 100}`

	// 创建结构体变量
	var book Book

	// 解析 JSON
	err := json.Unmarshal([]byte(jsonStr), &book)
	if err != nil {
		fmt.Println("JSON 解析错误:", err)
		return
	}

	// 输出解析后的数据
	fmt.Printf("书名: %s, 作者: %s, 价格: %d\n", book.Title, book.Author, book.Price)
}
