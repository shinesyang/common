package capture

// 这个程序只在 gameall里面调用,如何没有使用gameall 可以删除

import (
	"bufio"
	"os"

	"github.com/shinesyang/common"
	"go.uber.org/zap"
)

// 标记字符串,用来标记控制台日志捕捉完成(程序执行完成)

const MarkerString = "SfbRBRfi0VXvf0E7kcO5UBtuULZ9ibruvfVfT6EFM4XyRgj9BTYX4PCU5wH1HwNbsvWCwAqEnJsoyx9eech2g4jh2fTgdS4BhqEs"

var Logger *zap.SugaredLogger

// 捕捉link更新日志

type DataLogs interface {
	WriteStdout(b []byte)
	WriteStderr(b []byte)
}

func CaptureLogs(loads ...DataLogs) {
	stdout := os.Stdout
	stderr := os.Stderr
	r, w, _ := os.Pipe()
	er, ew, _ := os.Pipe()

	// 自定义logger
	stdout = w
	stderr = ew
	Logger = common.CustomLogger(stdout, stderr)

	// 捕捉stdout
	go func() {
		reader := bufio.NewReader(r)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				Logger.Errorf("读取输出时发生错误: %v", err)
				break
			}
			for _, load := range loads {
				load.WriteStdout(line)
			}
			os.Stdout.Write(line)
			os.Stdout.Write([]byte("\n"))
		}
	}()

	// 捕捉 stderr
	go func() {
		reader := bufio.NewReader(er)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				Logger.Errorf("读取输出时发生错误: %v", err)
				break
			}
			for _, load := range loads {
				load.WriteStderr(line)
			}
			os.Stdout.Write(line)
			os.Stdout.Write([]byte("\n"))
		}
	}()
}
