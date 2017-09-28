package main

import (
	"encoding/json"
	"fmt"
)

// Movie 电影
// 参考https://studygolang.com/articles/10001
type Movie struct {
	Title string `json:"title"`           // 序列化为title
	Year  int    `json:"-"`               // 不序列化
	Color bool   `json:"color,omitempty"` // 为空忽略
	Nice  bool   `json:",string"`         // 序列化为子符串
}

func main() {
	var movie = Movie{
		Title: "战狼2",
		Year:  2017,
		Color: false,
		Nice:  false,
	}
	data, err := json.Marshal(movie)
	if err != nil {
		fmt.Printf("失败：%s\n", err)
	} else {
		fmt.Printf("结果：%s\n", data)
	}

	data2, err := json.MarshalIndent(movie, "", "  ")
	if err != nil {
		fmt.Printf("失败：%s\n", err)
	} else {
		fmt.Printf("结果：%s\n", data2)
	}
}
