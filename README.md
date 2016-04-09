# Gebeta-logger

logger for [Gbeta](https://github.com/yyrdl/Gbeta)
![gbeta_logger_image](http://d.pcs.baidu.com/thumbnail/dd58087b8ab230c66aad36f073b083e0?fid=758558173-250528-38605564578761&time=1460178000&rt=sh&sign=FDTAER-DCb740ccc5511e5e8fedcff06b081203-H%2BvsboAu8XGGpJ43tL%2BfamnqA8Y%3D&expires=8h&chkv=0&chkbd=0&chkpc=&dp-logid=2308679241730188886&dp-callid=0&size=c710_u500&quality=100)
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