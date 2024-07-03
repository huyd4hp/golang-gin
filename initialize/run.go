package initialize

func Run(){
	InitMySQL()
	r := InitRouter()
	r.Run()
}