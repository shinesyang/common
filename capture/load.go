package capture

import (
	"fmt"
)

/*
	1. CaptureLogs 接收的参数是一个(DataLogs)接口类型数据
	2. 调用时只需要传入一个实现了WriteStdout/WriteStderr方法的接口即可
	3. 这里基于chan实现了一个 LoadLogs,
	4. 注意如果是使用其他方法处理数据,一定要保证使用的是线程安全数据类型,如: chan,sync.Map
	5. 或者使用锁来实现
*/

type LoadLogs struct {
	Logs   chan map[string]interface{}
	Marker string
}

func NewLoadLogs(logs chan map[string]interface{}, marker string) *LoadLogs {
	return &LoadLogs{
		logs,
		marker,
	}
}

func (s *LoadLogs) WriteStdout(b []byte) {
	go func() {
		s.Logs <- map[string]interface{}{s.Marker: string(b)}
	}()
}

func (s *LoadLogs) WriteStderr(b []byte) {
	go func() {
		msg := fmt.Sprintf(`<p1 style="color:#cdf602;">%s</p1>`, string(b))
		s.Logs <- map[string]interface{}{s.Marker: msg}
	}()
}
