# Gebeta-logger

logger for [Gbeta](https://github.com/yyrdl/Gbeta)


![gbeta_logger_image](http://d.pcs.baidu.com/thumbnail/f9c584af57c4d599ef3c4d07644104e3?fid=758558173-250528-157792661795900&time=1460178000&rt=sh&sign=FDTAER-DCb740ccc5511e5e8fedcff06b081203-QpbfgkCPK92Iwx2zrgvMck4A%2Fn4%3D&expires=8h&chkv=0&chkbd=0&chkpc=&dp-logid=2308795735709843142&dp-callid=0&size=c710_u500&quality=100)
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
