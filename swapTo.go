package common

import "encoding/json"

// struct 之间相关转换(借助于json)
func SwapTo(s1, s2 interface{}) error {
	// 数据先转成byte
	bytes, err := json.Marshal(s1)
	if err != nil {
		return err
	}
	// 再把数据转成新的struct
	return json.Unmarshal(bytes, s2)
}
