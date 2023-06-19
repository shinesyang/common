package common

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

/*
	解析yml文件
	filePath 文件路径
	data 解析之后的类型,可以是struct,可以是map,且必须是指针类型
*/

func ParseYml(filePath string, data interface{}) error {
	readFile, err := ioutil.ReadFile(filePath) // 读取文件
	if err != nil {
		return err
	}
	// 返回yaml
	return yaml.Unmarshal(readFile, data)
}
