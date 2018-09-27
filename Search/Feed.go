package Search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URL  string `json:"link"`
	Type string `json:"type"`
}

//读取并且反序列化源数据文件
func RetrieveFeeds() ([]*Feed, error) {
	//打开文件
	file, err := os.Open(dataFile)

	//当函数返回时 关闭文件
	defer  file.Close()

	if err != nil {
		return nil, err
	}

	//将文件解码到一个切片里
	//这个切片的每一项是指向一个Feed的指针
	var feeds []*Feed
	err=json.NewDecoder(file).Decode(&feeds)

	//这个函数不需要检查错误,调用者会做这件事
	return feeds,err
}
