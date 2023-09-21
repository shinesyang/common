# common

提交mysql prometheus NewTracer

```shell
conifg.go 定义consul配置中心
mysql.go 用来读取consul中的配置文件
jaeger.go 链路追踪
prometheus.go 监控
log.go 输入日志
capture 基于log.go实现的控制台日志捕捉程序
```

`capture在程序调用另一个程序时,不需修改被调用程序的内部代码,只需要再被调用的程序主入口处引入则可以完成被调用程序控制台日志捕捉`
```go
    /*capture使用方法*/
	func Run(logs chan map[string]interface{},marker string){
        defer func() {
            if err := recover(); err != nil {
            capture.Logger.Info(capture.MarkerString)
            }
        }()
        
        loadLogs := capture.NewLoadLogs(logs, marker)
		/*新增自定义捕捉到的日志重写到指定的日志文件内*/
        capture.CaptureLogs("capturelogs.log","capture_error.log",loadLogs)
        
        // 程序的最后****
        capture.Logger.Info(capture.MarkerString)
    }
```

`在NewLoadLogs是实现了一种方式,CaptureLogs只需要传入满足WriteStdout/WriteStderr的方法的对象即可,NewLoadLogs有说明`

#### 更多等待补充