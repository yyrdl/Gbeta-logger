# Gbeta-logger

logger for [Gbeta](https://github.com/yyrdl/Gbeta)


![gbeta_logger_image](http://studygolang.qiniudn.com/160409/ebec9fe8a5a973df2906de7a2e07cc01.png)
# Useage
```go
    import(
	  "github.com/yyrdl/gbeta"
	  "github.com/yyrdl/gbeta_logger"
	)
	func main(){
		app:=gbeta.App()
		app.WrapServeHTTP(gbeta_logger.Log)
		//.........
	}
```
