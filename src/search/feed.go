package search

import (
	"os"
	"encoding/json"
)

/**
 * 声明常量时，不需要指定类型
 */
const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// 当之前os.Open函数报错当时候，defer 保证 file.Close()
	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}